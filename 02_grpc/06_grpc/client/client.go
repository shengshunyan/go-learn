package main

import (
	"context"
	"fmt"
	proto2 "go-learn/02_grpc/06_grpc/proto"
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

	client := proto2.NewHelloWorldClient(conn)
	rsp, err := client.SayHello(context.Background(), &proto2.HelloRequest{
		Name: "shane",
	})
	if err != nil {
		panic("call SayHello failed" + err.Error())
	}
	fmt.Println(rsp.Message)
}
