package main

import (
	"encoding/json"
	"fmt"
	"go-learn/02_grpc/05_protobuf/helloworld"
	"google.golang.org/protobuf/proto"
)

type JsonHelloRequest struct {
	Name string `json:"name"`
}

func main() {
	req := helloworld.HelloRequest{
		Name: "shane",
	}
	// 序列化
	rsp, _ := proto.Marshal(&req)
	fmt.Println(string(rsp))
	fmt.Println(len(rsp))
	// 反序列化
	newReq := helloworld.HelloRequest{}
	_ = proto.Unmarshal(rsp, &newReq)
	fmt.Println(newReq.Name)

	// json对比
	req1 := JsonHelloRequest{
		Name: "shane",
	}
	rsp1, _ := json.Marshal(&req1)
	fmt.Println(string(rsp1))
	fmt.Println(len(rsp1))
}
