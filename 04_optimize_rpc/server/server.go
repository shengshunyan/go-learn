package main

import (
	"go-learn/04_optimize_rpc/handler"
	"go-learn/04_optimize_rpc/service_proxy"
	"net"
	"net/rpc"
)

func main() {
	// 1.实例化一个server
	listener, _ := net.Listen("tcp", ":1234")
	// 2. 注册处理逻辑 handler
	service_proxy.RegisterHelloService(&handler.HelloService{})
	// 3. 启动服务
	for {
		conn, _ := listener.Accept() // 当一个新的连接进来
		rpc.ServeConn(conn)
	}
}
