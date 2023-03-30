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

func main() {
	// 1.实例化一个server
	listener, _ := net.Listen("tcp", ":9000")
	// 2.注册处理逻辑
	rpc.RegisterName("rpcServer", &rpcServer{})

	//3. 启动服务
	for {
		conn, _ := listener.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
