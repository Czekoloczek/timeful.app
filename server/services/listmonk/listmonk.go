package listmonk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"schej.it/server/logger"
	"schej.it/server/utils"
)

// Adds the given user to the Listmonk contact list
// If subscriberId is not nil, then UPDATE the user instead of adding user
func AddUserToListmonk(email string, firstName string, lastName string, picture string, subscriberId *int, sendMarketingEmails bool) {
	// Listmonk is opt-in; only enable when explicitly set to "true".
	if os.Getenv("LISTMONK_ENABLED") != "true" {
		return
	}

	url := os.Getenv("LISTMONK_URL")
	username := os.Getenv("LISTMONK_USERNAME")
	password := os.Getenv("LISTMONK_PASSWORD")
	listIdString := os.Getenv("LISTMONK_LIST_ID")

	listId, err := strconv.Atoi(listIdString)
	if err != nil {
		logger.StdErr.Println(err)
		return
	}

	// Create new subscriber
	args := bson.M{
		"email":  email,
		"name":   firstName + " " + lastName,
		"status": "enabled",
		"attribs": bson.M{
			"firstName": firstName,
			"lastName":  lastName,
			"picture":   picture,
		},
		"preconfirm_subscriptions": true,
	}
	if sendMarketingEmails {
		args["lists"] = bson.A{listId}
	}
	body, _ := json.Marshal(args)
	bodyBuffer := bytes.NewBuffer(body)

	var req *http.Request
	if subscriberId != nil {
		// Existing subscriber
		req, _ = http.NewRequest("PUT", fmt.Sprintf("%s/api/subscribers/%d", url, *subscriberId), bodyBuffer)
	} else {
		// New subscriber
		req, _ = http.NewRequest("POST", fmt.Sprintf("%s/api/subscribers", url), bodyBuffer)
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.StdErr.Println(err)
		return
	}
	defer resp.Body.Close()
}

// Check if the user is already in listmonk
// Returns a bool representing whether the subscriber exists and the id of the subscriber if it does exist
func DoesUserExist(email string) (bool, *int) {
	// Listmonk is opt-in; only enable when explicitly set to "true".
	if os.Getenv("LISTMONK_ENABLED") != "true" {
		return false, nil
	}

	url := os.Getenv("LISTMONK_URL")
	username := os.Getenv("LISTMONK_USERNAME")
	password := os.Getenv("LISTMONK_PASSWORD")

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/subscribers?query=subscribers.email='%s'", url, email), nil)
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.StdErr.Println(err)
		return false, nil
	}
	defer resp.Body.Close()

	var response struct {
		Data struct {
			Results []struct {
				Id int `json:"id"`
			} `json:"results"`
		} `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		logger.StdErr.Println(err)
		return false, nil
	}

	if len(response.Data.Results) > 0 {
		return true, &response.Data.Results[0].Id
	} else {
		return false, nil
	}
}

// Send a transactional email using the specified template and data
func SendEmail(email string, templateId int, data bson.M) {
	// Listmonk is opt-in; only enable when explicitly set to "true".
	if os.Getenv("LISTMONK_ENABLED") != "true" {
		sendSMTPFallback(email, templateId, data)
		return
	}

	// Get listmonk url env vars
	listmonkUrl := os.Getenv("LISTMONK_URL")
	listmonkUsername := os.Getenv("LISTMONK_USERNAME")
	listmonkPassword := os.Getenv("LISTMONK_PASSWORD")

	// Construct body
	body, err := json.Marshal(bson.M{
		"subscriber_email": email,
		"template_id":      templateId,
		"data":             data,
		"content_type":     "html",
	})
	if err != nil {
		logger.StdErr.Println(err)
		return
	}

	// Construct request
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/tx", listmonkUrl), bytes.NewBuffer(body))
	req.SetBasicAuth(listmonkUsername, listmonkPassword)
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.StdErr.Println(err)
	}
	defer response.Body.Close()
}

// Send a transactional email using the specified template and data. Adds subscriber if they don't exist
func SendEmailAddSubscriberIfNotExist(email string, templateId int, data bson.M, sendMarketingEmails bool) {
	// Listmonk is opt-in; only enable when explicitly set to "true".
	if os.Getenv("LISTMONK_ENABLED") != "true" {
		sendSMTPFallback(email, templateId, data)
		return
	}

	if exists, _ := DoesUserExist(email); !exists {
		AddUserToListmonk(email, "", "", "", nil, sendMarketingEmails)
	}

	SendEmail(email, templateId, data)
}

func sendSMTPFallback(email string, templateId int, data bson.M) {
	subject, body := buildSMTPFallbackEmail(data)
	utils.SendEmail(email, subject, body, "text/plain")
}

func buildSMTPFallbackEmail(data bson.M) (string, string) {
	subject := "Timeful notification"
	if eventName, ok := data["eventName"].(string); ok && eventName != "" {
		subject = fmt.Sprintf("Timeful: %s", eventName)
	} else if groupName, ok := data["groupName"].(string); ok && groupName != "" {
		subject = fmt.Sprintf("Timeful: %s", groupName)
	}

	lines := []string{"Hello,"}
	if ownerName, ok := data["ownerName"].(string); ok && ownerName != "" {
		lines = append(lines, fmt.Sprintf("%s sent you a Timeful update.", ownerName))
	}
	if respondentName, ok := data["respondentName"].(string); ok && respondentName != "" {
		lines = append(lines, fmt.Sprintf("%s responded.", respondentName))
	}
	if numResponses, ok := data["numResponses"]; ok {
		lines = append(lines, fmt.Sprintf("Responses so far: %v", numResponses))
	}
	if eventName, ok := data["eventName"].(string); ok && eventName != "" {
		lines = append(lines, fmt.Sprintf("Event: %s", eventName))
	}
	if groupName, ok := data["groupName"].(string); ok && groupName != "" {
		lines = append(lines, fmt.Sprintf("Group: %s", groupName))
	}
	if eventUrl, ok := data["eventUrl"].(string); ok && eventUrl != "" {
		lines = append(lines, fmt.Sprintf("Open: %s", eventUrl))
	}
	if groupUrl, ok := data["groupUrl"].(string); ok && groupUrl != "" {
		lines = append(lines, fmt.Sprintf("Open: %s", groupUrl))
	}
	if finishedUrl, ok := data["finishedUrl"].(string); ok && finishedUrl != "" {
		lines = append(lines, fmt.Sprintf("Already responded: %s", finishedUrl))
	}
	if emails, ok := data["emails"]; ok {
		lines = append(lines, fmt.Sprintf("New attendees: %v", emails))
	}

	return subject, strings.Join(lines, "\n")
}
