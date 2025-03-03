package main

import (
	"fmt"
	"go-learn/02_grpc/11_grpc_validate/proto"
)

func main() {
	person := proto.Person{
		Id: 1000,
		//Email: "shane@163.com",
		//Name:  "shane",
		Home: &proto.Person_Location{
			Lat: 12,
			Lng: 34,
		},
	}

	// 检查第一个错误
	err := person.Validate()
	if err != nil {
		fmt.Println("validate error", err.Error())
	}

	// 检查所有错误
	err1 := person.ValidateAll()
	if err1 != nil {
		fmt.Println("validate error", err1.Error())
	}
}
