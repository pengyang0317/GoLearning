package main

import (
	"context"
	"net"

	hellopb "lgo/src/ch4/v2/proto"

	"google.golang.org/grpc"
)

type Server struct {
	hellopb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *hellopb.HelloRequesto) (*hellopb.HelloResponse, error) {
	return &hellopb.HelloResponse{
		Message: "Hello " + in.Name,
	}, nil
}

func main() {

	g := grpc.NewServer()

	hellopb.RegisterGreeterServer(g, &Server{})

	// 要查看 80 端口是否被占用，可以在终端中执行以下命令：
	// netstat -an | grep 8888
	// 该命令会列出所有的网络连接信息，并通过 grep 命令对其进行过滤，只显示与 80 端口相关的连接信息。如果输出结果中存在 LISTEN，则表示该端口正在被监听，即该端口已被占用。如果输出结果为空，则表示该端口未被占用。
	Listener, err := net.Listen("tcp", ":8888")

	if err != nil {
		panic(err)
	}

	err = g.Serve(Listener)
	if err != nil {
		panic(err)
	}

}
