package main

import (
	"context"
	"fmt"
	proto2 "go-learn/02_grpc/09_grpc_metadata/proto"
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

	client := proto2.NewHelloWorldClient(conn)
	md := metadata.Pairs("name", "secret")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	rsp, err := client.SayHello(ctx, &proto2.HelloRequest{
		Name: "shane",
	})
	if err != nil {
		panic("call SayHello failed" + err.Error())
	}
	fmt.Println(rsp.Message)
}
