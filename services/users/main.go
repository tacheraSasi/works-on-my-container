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

	return &pb.SendEmailResponse{
		Message: "Email is being sent ",
	}, nil
}

func main() {
	listener, err := shared.StartGRPCServer("8081")
	if err != nil {
		log.Fatal(err)
	}
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "redis:6379"})
	defer client.Close()

	// Enqueue immediately
	task, _ := tasks.NewWelcomeEmailTask("123", "tach@example.com")
	info, err := client.Enqueue(task)
	println("Enqueued task with ID:", info.ID)
	// info.ID gives you the task ID

	// Enqueue with a delay (process in 5 minutes)
	info, err = client.Enqueue(task, asynq.ProcessIn(5*time.Minute))

	// Enqueue with max retry
	info, err = client.Enqueue(task, asynq.MaxRetry(3))

	// Enqueue to a specific queue with priority
	info, err = client.Enqueue(task, asynq.Queue("critical"))

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})

	fmt.Println("Server running on :8081")
	grpcServer.Serve(listener)
}
