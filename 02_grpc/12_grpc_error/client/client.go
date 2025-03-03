package main

import (
	"context"
	"fmt"
	proto2 "go-learn/02_grpc/12_grpc_error/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
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
	// 客户端解析grpc返回的错误
	if err != nil {
		//panic("call SayHello failed" + err.Error())
		st, ok := status.FromError(err)
		if !ok {
			panic("解析error失败")
		}
		fmt.Println(st.Code(), st.Message())
	} else {
		fmt.Println(rsp.Message)
	}
}
