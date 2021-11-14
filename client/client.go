package main

import (
	"context"
	"github.com/riverphillips/go-grpc/client/message"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	helloWorldServiceClient := message.NewHelloWorldServiceClient(conn)

	response, err := helloWorldServiceClient.HelloWorld(context.Background(), &message.HelloWorldInput{Name: "River"})
	if err != nil {
		log.Fatalf("Error when calling hello world: %v", err)
	}

	log.Printf("Response from server: %s", response.Greeting)

}
