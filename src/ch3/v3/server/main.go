package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type rpcServer struct{}

func (s *rpcServer) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}

func main() {
	err := rpc.RegisterName("rpcServer", new(rpcServer))
	if err != nil {
		fmt.Println("register error:", err)
	}

	http.HandleFunc("/httprpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":9000", nil)
}
