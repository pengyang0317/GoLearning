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

const PROT = ":9000"

func main() {
	//注册服务：使用 rpc.RegisterName 方法注册一个名为 "rpcServer" 的服务，并将其关联到一个实现了对应方法的结构体 rpcServer 上。
	err := rpc.RegisterName("rpcServer", new(rpcServer))
	if err != nil {
		fmt.Println("register error:", err)
	}
	//处理 HTTP 请求：使用 http.HandleFunc 方法为路径 "/httprpc" 注册一个处理函数，该函数用于处理 HTTP 请求。在处理函数中，创建一个 io.ReadWriteCloser 对象，将 HTTP 请求的 Body 和 ResponseWriter 封装起来，并使用 rpc.ServeRequest 方法将请求交给 RPC 服务端处理。
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
	//启动 HTTP 服务：使用 http.ListenAndServe 方法启动一个 HTTP 服务，监听指定的端口（PROT），等待客户端的请求。
	http.ListenAndServe(PROT, nil)
}
