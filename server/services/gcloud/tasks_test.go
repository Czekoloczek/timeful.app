package gcloud

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"schej.it/server/logger"
)

func TestCreateEmailTask(t *testing.T) {
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
		t.Skip("Skipping gcloud tasks test; .env not configured")
	}

	InitTasks()
	CreateEmailTask("schej.team@gmail.com", "Jonathan", "casablanca", "65e636bb760d3ea2e113e161")
}

func TestDeleteEmailTask(t *testing.T) {
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
		t.Skip("Skipping gcloud tasks test; .env not configured")
	}

	InitTasks()

	// Should fail
	fmt.Println("Delete email task that doesn't exist...")
	DeleteEmailTask("id_that_doesn't_exist")
	fmt.Println("Should have thrown an error ^")

	// Should succeed
	fmt.Println("Creating email task...")
	taskIds := CreateEmailTask("schej.team@gmail.com", "Jonathan", "casablanca", "65e636bb760d3ea2e113e161")
	fmt.Println("Email task created")

	time.Sleep(10 * time.Second)
	for _, taskId := range taskIds {
		fmt.Println("Deleting email task with taskId: ", taskId)
		DeleteEmailTask(taskId)
		fmt.Println("Deleted email task with taskId: ", taskId)
	}

	fmt.Println("Done.")
}

func TestCreateLocalEmailTasks(t *testing.T) {
	var called int
	originalSender := sendReminderEmailSMTPFunc
	sendReminderEmailSMTPFunc = func(email string, ownerName string, eventName string, eventUrl string, finishedUrl string, label string) {
		called++
	}
	defer func() { sendReminderEmailSMTPFunc = originalSender }()

	taskIds := createLocalEmailTasks("test@example.com", "Ada", "Demo", "event123")
	if len(taskIds) != 3 {
		t.Fatalf("expected 3 task ids, got %d", len(taskIds))
	}

	for i := 0; i < 120 && called == 0; i++ {
		time.Sleep(10 * time.Millisecond)
	}

	if called == 0 {
		t.Fatalf("expected local reminder sender to be called at least once")
	}

	for _, id := range taskIds {
		if _, ok := localTasks.Load(id); !ok {
			t.Fatalf("expected task %s to be stored", id)
		}
		DeleteEmailTask(id)
	}
}
