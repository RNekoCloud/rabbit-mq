package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/AvinFajarF/proto"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	rabbitMQURL = "amqp://guest:guest@localhost:5672/"
)

type server struct {
	pb.UnimplementedMessageServiceServer
}

func (s *server) ConsumeMessages(req *pb.ConsumeRequest, stream pb.MessageService_ConsumeMessagesServer) error {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return status.Errorf(codes.Internal, "failed to open RabbitMQ channel: %v", err)
	}
	defer ch.Close()

	req.QueueName = "TestQueue"
	msgs, err := ch.Consume(
		req.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to consume messages: %v", err)
	}

	for d := range msgs {
		if err := stream.Send(&pb.Message{Body: d.Body}); err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}

	return nil
}

func main() {
	fmt.Println("gRPC RabbitMQ Server")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMessageServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
