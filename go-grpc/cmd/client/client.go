package main

import (
	"context"
	"fmt"
	"go-grpc/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	connection, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	// AddUser(client)
	// AddUserVerbose(client)
	// AddUsers(client)
	AddUserStreamBoth(client)
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

func AddUsers(client pb.UserServiceClient) {
	req := []*pb.User{
		&pb.User{
			Id:    "1",
			Name:  "John Doe",
			Email: "john_doe@gmail.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "John Doe2",
			Email: "john_doe@gmail.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "John Doe3",
			Email: "john_doe@gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Could not open stream: %v", err)
	}

	for _, user := range req {
		stream.Send(user)
		time.Sleep(time.Second * 2)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Could not receive response: %v", err)
	}

	fmt.Println(res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {
	stream, err := client.AddUserStreamBoth(context.Background())

	if err != nil {
		log.Fatalf("Could not open stream: %v", err)
	}

	req := []*pb.User{
		&pb.User{
			Id:    "1",
			Name:  "John Doe",
			Email: "john_doe@gmail.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "John Doe2",
			Email: "john_doe@gmail.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "John Doe3",
			Email: "john_doe@gmail.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, user := range req {
			fmt.Println("Sending:", user.Name)
			stream.Send(user)
			time.Sleep(time.Second * 2)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Could not receive stream: %v", err)
			}

			fmt.Println("Status:", res.Status)
		}
		close(wait)
	}()

	<-wait
}
