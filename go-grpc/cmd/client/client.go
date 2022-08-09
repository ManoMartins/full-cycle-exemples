package main

import (
	"context"
	"fmt"
	"go-grpc/pb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	connection, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	// AddUser(client)
	AddUserVerbose(client)
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

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "1",
		Name:  "John Doe",
		Email: "john_doe@hmailc.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		panic(err)
	}

	for {
		stream, err := responseStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receive stream: %v", err)
		}

		fmt.Println("Status:", stream.Status)

	}
}
