package main

import (
	"context"
	"fmt"
	"go-learn/09_grpc_metadata/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.NewClient("localhost:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("new client failed" + err.Error())
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic("close client failed" + err.Error())
		}
	}(conn)

	client := proto.NewHelloWorldClient(conn)
	md := metadata.Pairs("name", "secret")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	rsp, err := client.SayHello(ctx, &proto.HelloRequest{
		Name: "shane",
	})
	if err != nil {
		panic("call SayHello failed" + err.Error())
	}
	fmt.Println(rsp.Message)
}
