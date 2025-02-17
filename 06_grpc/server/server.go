package main

import (
	"context"
	"go-learn/06_grpc/proto"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	proto.UnimplementedHelloWorldServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{Message: "Hello " + request.Name}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterHelloWorldServer(g, &Server{})
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	err = g.Serve(lis)
	if err != nil {
		panic("failed to serve: " + err.Error())
	}
}
