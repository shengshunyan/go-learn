package client_proxy

import (
	"go-learn/02_grpc/04_optimize_rpc/handler"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(network, address string) *HelloServiceStub {
	client, err := rpc.Dial(network, address)
	if err != nil {
		panic("create client error")
	}
	return &HelloServiceStub{client}
}

func (c *HelloServiceStub) Hello(args string, reply *string) error {
	err := c.Call(handler.HelloServiceName+".Hello", args, reply)
	if err != nil {
		return err
	}
	return nil
}
