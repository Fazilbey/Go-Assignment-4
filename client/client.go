package main

import (
	"assignment-4/protos"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"log"
)

func main() {
	connectiion, error := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if error != nil {
		log.Fatalf("Failed to connect: %v", error)
	}
	defer connectiion.Close()

	clientdata := protos.NewUserServiceClient(connectiion)
	addUserResp, error := clientdata.AddUser(context.Background(), &protos.UserCreateRequest{Name: "Fazil", Email: "fazil@example.com"})
	if error != nil {
		log.Fatalf("AddUser call failed: %v", error)
	}
	fmt.Printf("User ID: %d\n", addUserResp.Id)
	getUserResp, error := clientdata.GetUser(context.Background(), &protos.UserGetRequest{Id: 1})

	if error != nil {
		log.Fatalf("GetUser call failed: %v", error)
	}
	fmt.Printf("User: %+v\n", getUserResp)
	listUsersResp, error := clientdata.ListUsers(context.Background(), &protos.UserGetAllRequest{})

	if error != nil {
		log.Fatalf("ListUsers call failed: %v", error)
	}
	fmt.Println("Users:")
	for _, user := range listUsersResp.Users {
		fmt.Printf(" - %+v\n", user)
	}
}
