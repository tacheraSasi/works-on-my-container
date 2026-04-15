package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	pb "works-on-my-machine/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type clients struct {
	userService pb.UserServiceClient
}

func connectToClients(clients *clients) (*clients, error) {
	return clients, nil
}

func handler(usersClient pb.UserServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		resp, err := usersClient.GetUser(ctx, &pb.GetUserRequest{Id: "123"})
		if err != nil {
			fmt.Fprintf(w, "Error calling GetUser: %v", err)
			return
		}

		fmt.Fprintf(w, "Hello from gateway! User: %s", resp.Name)
	}

}

func main() {
	conn, err := grpc.NewClient("users-service:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println("Failed to connect to users service:", err)
		return
	}
	defer conn.Close()

	usersServiceClient := pb.NewUserServiceClient(conn)

	clients := &clients{
		userService: usersServiceClient,
	}

	c, err := connectToClients(clients)
	if err != nil {
		fmt.Println("Failed to connect to clients:", err)
		return
	}
	usersClient := c.userService

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler(usersClient))

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", mux)
}
