package listmonk

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"schej.it/server/logger"
)

func TestSendEmail(t *testing.T) {
	// Init logfile
	logFile, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Init logger
	logger.Init(logFile)

	// Load .env file
	err = godotenv.Load("../../.env")
	if err != nil {
		t.Skip("Skipping listmonk email test; .env not configured")
	}

	SendEmail("schej.team@gmail.com", 8, bson.M{
		"eventName": "casablanca",
		"eventUrl":  "http://localhost:8080/e/65e636bb760d3ea2e113e161",
	})
}

func TestBuildSMTPFallbackEmail(t *testing.T) {
	subject, body := buildSMTPFallbackEmail(bson.M{
		"eventName":      "Demo Event",
		"ownerName":      "Ada",
		"respondentName": "Max",
		"eventUrl":       "https://example.com/e/123",
	})

	if !strings.Contains(subject, "Demo Event") {
		t.Fatalf("expected subject to include event name, got %q", subject)
	}
	if !strings.Contains(body, "Ada") || !strings.Contains(body, "Max") {
		t.Fatalf("expected body to include owner/respondent, got %q", body)
	}
	if !strings.Contains(body, "https://example.com/e/123") {
		t.Fatalf("expected body to include event url, got %q", body)
	}
}
