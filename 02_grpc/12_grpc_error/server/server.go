package main

import (
	"context"
	proto2 "go-learn/02_grpc/12_grpc_error/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type Server struct {
	proto2.UnimplementedHelloWorldServer
}

func (s *Server) SayHello(ctx context.Context, request *proto2.HelloRequest) (*proto2.HelloResponse, error) {
	// 返回grpc错误
	return nil, status.Error(codes.NotFound, "记录没有找到")
	//return &proto.HelloResponse{Message: "Hello " + request.Name}, nil
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
