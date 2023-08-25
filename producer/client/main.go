package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/AvinFajarF/proto"
	"google.golang.org/grpc"
)

const (
	serverAddress = "localhost:50052"
)

func main() {
	fmt.Println("gRPC RabbitMQ Client")

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMessageServiceClient(conn)

	message := &pb.Message{
		Body:        []byte("Hello from client"),
		ContentType: "text/plain",
	}

	response, err := client.PublishMessage(context.Background(), message)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	fmt.Printf("Publish Response: Success=%v, Message=%s\n", response.Success, response.Message)
}
