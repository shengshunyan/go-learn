package main

import (
	"context"
	"fmt"
	proto2 "go-learn/02_grpc/07_stream_grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("new client failed" + err.Error())
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic("close client failed" + err.Error())
		}
	}(conn)
	client := proto2.NewGreeterClient(conn)

	// 服务端流模式
	//rsp, err := client.ServerStream(context.Background(), &proto.StreamReq{
	//	Data: "hi, shane!",
	//})
	//for {
	//	a, err := rsp.Recv()
	//	if err != nil {
	//		fmt.Println("server receive failed" + err.Error())
	//		break
	//	}
	//	fmt.Println(a)
	//}

	// 客户端流模式
	//stream, err := client.ClientStream(context.Background())
	//if err != nil {
	//	panic("client stream failed" + err.Error())
	//}
	//i := 0
	//for {
	//	i++
	//	_ = stream.Send(&proto.StreamReq{
	//		Data: fmt.Sprintf("%v", time.Now().Unix()),
	//	})
	//	time.Sleep(time.Second)
	//	if i > 10 {
	//		break
	//	}
	//}

	// 双向流模式
	stream, err := client.AllStream(context.Background())
	if err != nil {
		panic("client stream failed" + err.Error())
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		i := 0
		for {
			i++
			_ = stream.Send(&proto2.StreamReq{
				Data: fmt.Sprintf("client send data: %v", time.Now().Unix()),
			})
			time.Sleep(time.Second)
			if i > 10 {
				break
			}
		}
	}()

	go func() {
		defer wg.Done()

		for {
			in, err := stream.Recv()
			if err != nil {
				fmt.Println("client receive failed" + err.Error())
				break
			}
			fmt.Println("client receive data:", in.Data)
		}
	}()

	wg.Wait()

}
