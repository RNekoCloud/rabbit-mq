// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/email.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageServiceClient interface {
	ConsumeMessages(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (MessageService_ConsumeMessagesClient, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) ConsumeMessages(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (MessageService_ConsumeMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &MessageService_ServiceDesc.Streams[0], "/proto.MessageService/ConsumeMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageServiceConsumeMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MessageService_ConsumeMessagesClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type messageServiceConsumeMessagesClient struct {
	grpc.ClientStream
}

func (x *messageServiceConsumeMessagesClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessageServiceServer is the server API for MessageService service.
// All implementations must embed UnimplementedMessageServiceServer
// for forward compatibility
type MessageServiceServer interface {
	ConsumeMessages(*ConsumeRequest, MessageService_ConsumeMessagesServer) error
	mustEmbedUnimplementedMessageServiceServer()
}

// UnimplementedMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (UnimplementedMessageServiceServer) ConsumeMessages(*ConsumeRequest, MessageService_ConsumeMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method ConsumeMessages not implemented")
}
func (UnimplementedMessageServiceServer) mustEmbedUnimplementedMessageServiceServer() {}

// UnsafeMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServiceServer will
// result in compilation errors.
type UnsafeMessageServiceServer interface {
	mustEmbedUnimplementedMessageServiceServer()
}

func RegisterMessageServiceServer(s grpc.ServiceRegistrar, srv MessageServiceServer) {
	s.RegisterService(&MessageService_ServiceDesc, srv)
}

func _MessageService_ConsumeMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsumeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessageServiceServer).ConsumeMessages(m, &messageServiceConsumeMessagesServer{stream})
}

type MessageService_ConsumeMessagesServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type messageServiceConsumeMessagesServer struct {
	grpc.ServerStream
}

func (x *messageServiceConsumeMessagesServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

// MessageService_ServiceDesc is the grpc.ServiceDesc for MessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConsumeMessages",
			Handler:       _MessageService_ConsumeMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/email.proto",
}
