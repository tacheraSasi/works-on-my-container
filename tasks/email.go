package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const TypeWelcomeEmail = "email:welcome"

type WelcomeEmailPayload struct {
	UserID string
	Email  string
}

// NewWelcomeEmailTask creates the task (producer side)
func NewWelcomeEmailTask(userID, email string) (*asynq.Task, error) {
	payload, err := json.Marshal(WelcomeEmailPayload{UserID: userID, Email: email})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeWelcomeEmail, payload), nil
}
