package main

import (
	"context"
	"fmt"
	"log"
	"time"
	pb "works-on-my-machine/proto/user"
	"works-on-my-machine/shared"
	"works-on-my-machine/tasks"

	"github.com/hibiken/asynq"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
	asynqClient *asynq.Client
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Println("Received request for user:", req.Id)

	return &pb.GetUserResponse{
		Id:   req.Id,
		Name: "Tach",
	}, nil
}

func (s *server) SendEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	log.Println("Received request to send email to:", req.To)

	msg, err := sendEmailTask(s.asynqClient)
	if err != nil {
		return nil, fmt.Errorf("failed to enqueue email task: %v", err)
	}
	return &pb.SendEmailResponse{
		Message: msg,
	}, nil
}

func main() {
	listener, err := shared.StartGRPCServer("8081")
	if err != nil {
		log.Fatal(err)
	}
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "redis:6379"})
	defer client.Close()

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{asynqClient: client})

	fmt.Println("Server running on :8081")
	grpcServer.Serve(listener)
}

func sendEmailTask(client *asynq.Client) (string, error) {
	// Enqueue immediately
	task, _ := tasks.NewWelcomeEmailTask("123", "tach@example.com")
	info, err := client.Enqueue(task)
	if err != nil {
		return "", err
	}
	println("Enqueued task with ID:", info.ID)
	// info.ID gives you the task ID

	// Enqueue with a delay (process in 5 minutes)
	info, err = client.Enqueue(task, asynq.ProcessIn(5*time.Minute))
	if err != nil {
		return "", err
	}

	// Enqueue with max retry
	info, err = client.Enqueue(task, asynq.MaxRetry(3))
	if err != nil {
		return "", err
	}
	// Enqueue to a specific queue with priority
	info, err = client.Enqueue(task, asynq.Queue("critical"))
	if err != nil {
		return "", err
	}
	return "Email task enqueued with ID: " + info.ID, nil

}
