package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	pb "github.com/AvinFajarF/proto"
)

const (
	rabbitMQURL = "amqp://guest:guest@localhost:5672/"
)

type server struct{
	pb.UnimplementedMessageServiceServer
}

func (s *server) PublishMessage(ctx context.Context, msg *pb.Message) (*pb.PublishResponse, error) {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return &pb.PublishResponse{Success: false, Message: fmt.Sprintf("Failed to connect to RabbitMQ: %v", err)}, nil
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return &pb.PublishResponse{Success: false, Message: fmt.Sprintf("Failed to open RabbitMQ channel: %v", err)}, nil
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return &pb.PublishResponse{Success: false, Message: fmt.Sprintf("Failed to declare queue: %v", err)}, nil
	}

	msg.Body = []byte("Hello World")

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			Body:        msg.Body,
			ContentType: msg.ContentType,
		},
	)

	log.Println(string(msg.Body))
	if err != nil {
		return &pb.PublishResponse{Success: false, Message: fmt.Sprintf("Failed to publish message: %v", err)}, nil
	}

	return &pb.PublishResponse{Success: true, Message: "Message published successfully"}, nil
}

func main() {
	fmt.Println("gRPC RabbitMQ Publisher")

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMessageServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
