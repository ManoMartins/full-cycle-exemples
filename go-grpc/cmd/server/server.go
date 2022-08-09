package main

import (
	"go-grpc/pb"
	"go-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
