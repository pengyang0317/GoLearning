package main

import (
	"net"
	"net/rpc"
)

const PORT = ":9000"

type rpcServer struct{}

func (s *rpcServer) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}

func main() {
	// 1.实例化一个 server：使用 net.Listen 方法监听指定端口（PORT），创建一个 TCP 的 listener 对象，用于接收客户端的连接请求。
	listener, _ := net.Listen("tcp", PORT)
	// 2.注册处理逻辑：使用 rpc.RegisterName 方法注册一个名为 "rpcServer" 的服务，并将其关联到一个实现了对应方法的结构体 rpcServer 上。
	rpc.RegisterName("rpcServer", &rpcServer{})

	//3.启动服务：调用 listener.Accept() 方法接受客户端的连接请求，并通过 rpc.ServeConn 方法将客户端的请求交给 RPC 服务端处理。
	conn, _ := listener.Accept()
	rpc.ServeConn(conn)
}
