package main

import (
	"assignment-4/protos"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserServiceServerImplementation struct {
	protos.UnimplementedUserServiceServer
}

func (s *UserServiceServerImplementation) AddUser(ctx context.Context, req *protos.UserCreateRequest) (*protos.UserIdResponse, error) {
	fmt.Printf("Received request to add user with name: %s, email: %s\n", req.Name, req.Email)

	ID := int32(123)

	return &protos.UserIdResponse{Id: ID}, nil
}

func (s *UserServiceServerImplementation) GetUser(ctx context.Context, req *protos.UserGetRequest) (*protos.User, error) {
	return &protos.User{
		Id:    req.Id,
		Name:  "Fazil Bey",
		Email: "fazil@gmail.com",
	}, nil
}

func (s *UserServiceServerImplementation) ListUsers(ctx context.Context, req *protos.UserGetAllRequest) (*protos.UserGetAllResponse, error) {

	users := []*protos.User{
		{Id: 1, Name: "Gosha", Email: "Dudar@example.com"},
		{Id: 2, Name: "Bob", Email: "gren@example.com"},
		{Id: 3, Name: "Charlie", Email: "path@example.com"},
	}

	return &protos.UserGetAllResponse{Users: users}, nil
}

func main() {
	server_link := grpc.NewServer()

	protos.RegisterUserServiceServer(server_link, &UserServiceServerImplementation{})

	listener, err := net.Listen("tcp", ":50001")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Server started, poert 50001")
	if err := server_link.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
