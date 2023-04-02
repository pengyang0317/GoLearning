package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type rpcServer struct{}

func (s *rpcServer) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}

const PORT = ":9000"

func main() {
	// 1.实例化一个 server：使用 net.Listen 方法监听指定端口 9000，创建一个 TCP 的 listener 对象，用于接收客户端的连接请求。
	listener, _ := net.Listen("tcp", PORT)
	// 2.注册处理逻辑：使用 rpc.RegisterName 方法注册一个名为 "rpcServer" 的服务，并将其关联到一个实现了对应方法的结构体 rpcServer 上
	rpc.RegisterName("rpcServer", &rpcServer{})

	//3.启动服务：使用 listener.Accept() 方法阻塞等待客户端的连接请求，然后通过 jsonrpc.NewServerCodec 方法创建一个 JSON-RPC 编码的服务器编解码器，并通过 rpc.ServeCodec 方法将客户端的请求交给 RPC 服务端处理。由于 listener.Accept() 方法是阻塞的，需要使用一个 for 循环不断接收客户端的连接请求。
	for {
		conn, _ := listener.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
