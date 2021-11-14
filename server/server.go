package main

import (
	"fmt"
	"github.com/riverphillips/go-grpc/server/message"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	helloWorldServiceServer := message.Server{}

	grpcServer := grpc.NewServer()

	message.RegisterHelloWorldServiceServer(grpcServer, &helloWorldServiceServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
