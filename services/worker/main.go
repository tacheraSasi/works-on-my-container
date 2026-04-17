package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"works-on-my-machine/tasks"

	"github.com/hibiken/asynq"
)

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "redis:6379"},
		asynq.Config{
			Concurrency: 10, // processing 10 tasks at a time
			Queues: map[string]int{
				"critical": 6, // higher priority
				"default":  3,
				"low":      1,
			},
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc("email:send", func(ctx context.Context, t *asynq.Task) error {
		fmt.Printf("Received task: type=%s payload=%s\n", t.Type(), string(t.Payload()))
		var p tasks.WelcomeEmailPayload
		if err := json.Unmarshal(t.Payload(), &p); err != nil {
			return fmt.Errorf("json.Unmarshal failed: %v", err)
		}

		// Simulating Doing the actual work
		log.Printf("Sending welcome email to %s (user %s)", p.Email, p.UserID)
		time.Sleep(5 * time.Second)
		log.Printf("Welcome email sent to %s (user %s)", p.Email, p.UserID)

		// Return nil = success, return error = retry
		return nil
	})

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
