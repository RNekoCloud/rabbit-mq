package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/AvinFajarF/proto" // Import the generated proto package

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("gRPC RabbitMQ Client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMessageServiceClient(conn)

	req := &pb.ConsumeRequest{
		QueueName: "TestQueue", // Set the queue name you want to consume from
	}

	stream, err := client.ConsumeMessages(context.Background(), req)
	if err != nil {
		log.Fatalf("Error consuming messages: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving message: %v", err)
		}
		fmt.Printf("Received message: %s\n", msg.Body)
	}
}
