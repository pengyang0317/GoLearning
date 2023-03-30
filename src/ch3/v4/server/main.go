package main

import (
	"fmt"
	"lgo/src/ch3/v4/handler"
	sproxy "lgo/src/ch3/v4/s_proxy"
	"net"
	"net/rpc"
)

func main() {
	// 1.实例化一个server
	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		fmt.Println("listen error:", err)
	}
	// 2.注册处理逻辑
	sproxy.RegisterRpcServer(new(handler.RpcServer))
	//3. 启动服务
	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}

}
