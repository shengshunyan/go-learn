package main

import (
	"fmt"
	"go-learn/02_grpc/04_optimize_rpc/client_proxy"
)

func main() {
	// 1. 建立连接
	client := client_proxy.NewHelloServiceClient("tcp", "127.0.0.1:1234")

	// 2. 调用
	var reply string
	_ = client.Hello("shane", &reply)
	fmt.Println(reply)
}
