package main

import (
	"context"
	"fmt"
	"go-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	AddUser(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "1",
		Name:  "John Doe",
		Email: "john_doe@hmailc.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
