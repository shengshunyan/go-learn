package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello " + request
	return nil
}

func main() {
	// 1.实例化一个server
	_ = rpc.RegisterName("HelloService", &HelloService{})
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn = struct {
			io.Writer
			io.ReadCloser
		}{
			w,
			r.Body,
		}
		_ = rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	_ = http.ListenAndServe(":1234", nil)
}
