package service_proxy

import (
	"go-learn/02_grpc/04_optimize_rpc/handler"
	"net/rpc"
)

type IHelloService interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv IHelloService) {
	_ = rpc.RegisterName(handler.HelloServiceName, srv)
}
