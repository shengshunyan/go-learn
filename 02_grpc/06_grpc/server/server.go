package main

import (
	"context"
	proto2 "go-learn/02_grpc/06_grpc/proto"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	proto2.UnimplementedHelloWorldServer
}

func (s *Server) SayHello(ctx context.Context, request *proto2.HelloRequest) (*proto2.HelloResponse, error) {
	return &proto2.HelloResponse{Message: "Hello " + request.Name}, nil
}

func main() {
	g := grpc.NewServer()
	proto2.RegisterHelloWorldServer(g, &Server{})
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	err = g.Serve(lis)
	if err != nil {
		panic("failed to serve: " + err.Error())
	}
}
