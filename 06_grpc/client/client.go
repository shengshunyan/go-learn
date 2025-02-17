package main

import (
	"context"
	"fmt"
	"go-learn/06_grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	rsp, err := client.SayHello(context.Background(), &proto.HelloRequest{
		Name: "shane",
	})
	if err != nil {
		panic("call SayHello failed" + err.Error())
	}
	fmt.Println(rsp.Message)
}
