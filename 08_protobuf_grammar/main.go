package main

import (
	"fmt"
	"go-learn/08_protobuf_grammar/proto"
)

func main() {
	req := proto.HelloRequest{
		Name:   "shane",
		Age:    "18",
		Gender: proto.Gender_MALE,
		Mp: map[string]string{
			"num": "4",
		},
	}

	fmt.Println(&req)
}
