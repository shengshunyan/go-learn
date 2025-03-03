package main

import (
	"fmt"
	proto2 "go-learn/02_grpc/07_stream_grpc/proto"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

type Server struct {
	proto2.UnimplementedGreeterServer
}

func (s *Server) ServerStream(req *proto2.StreamReq, server proto2.Greeter_ServerStreamServer) error {
	i := 0
	for {
		i++
		_ = server.Send(&proto2.StreamResp{
			Data: fmt.Sprintf("[%v]: %s", time.Now().Unix(), req.Data),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	return nil
}

func (s *Server) ClientStream(server proto2.Greeter_ClientStreamServer) error {
	for {
		in, err := server.Recv()
		if err != nil {
			fmt.Println("server receive err:", err)
			break
		}
		fmt.Println("server receive data:", in.Data)
	}

	return nil
}

func (s *Server) AllStream(server proto2.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		for {
			in, err := server.Recv()
			if err != nil {
				fmt.Println("server receive err:", err)
				break
			}
			fmt.Println("server receive data:", in.Data)
		}
	}()

	go func() {
		defer wg.Done()

		i := 0
		for {
			i++
			_ = server.Send(&proto2.StreamResp{
				Data: fmt.Sprintf("server send data: %v", time.Now().Unix()),
			})
			time.Sleep(time.Second)
			if i > 10 {
				break
			}
		}
	}()

	wg.Wait()

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	g := grpc.NewServer()
	proto2.RegisterGreeterServer(g, &Server{})

	err = g.Serve(lis)
	if err != nil {
		panic("failed to serve: " + err.Error())
	}
}
