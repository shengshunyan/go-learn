package main

import (
	"context"
	"fmt"
	"go-learn/09_grpc_metadata/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
)

type Server struct {
	proto.UnimplementedHelloWorldServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("get md error")
	}
	//for k, v := range md {
	//	fmt.Println(k, v)
	//}
	if name, ok := md["name"]; ok {
		fmt.Println(name)
	}
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
