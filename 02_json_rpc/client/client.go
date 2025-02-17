package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 1. 建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		panic("connect error")
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	// 2. 调用
	var reply string
	err = client.Call("HelloService.Hello", "shane", &reply)
	if err != nil {
		panic("call error")
	}
	fmt.Println(reply)
}
