package main

import (
	"context"
	"fmt"
	"log"
	"works-on-my-machine/proto/user"
	"works-on-my-machine/shared"

	"google.golang.org/grpc"
)

type server struct {
	user.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	log.Println("Received request for user:", req.Id)

	return &user.GetUserResponse{
		Id:   req.Id,
		Name: "Tach",
	}, nil
}

func main() {
	listener, err := shared.StartGRPCServer("8081")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &server{})

	fmt.Println("Server running on :8081")
	grpcServer.Serve(listener)
}
