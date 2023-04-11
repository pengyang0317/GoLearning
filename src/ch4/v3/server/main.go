package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	streampb "lgo/src/ch4/v3/proto"

	"google.golang.org/grpc"
)

type Server struct {
	streampb.UnimplementedGreeterServer
}

// 服务端流模式
func (s *Server) ServerStream(req *streampb.StreamRequest, res streampb.Greeter_ServerStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&streampb.StreamResponse{
			Data: fmt.Sprintf("%v", req),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

// 客户端流模式
func (s *Server) ClientStream(res streampb.Greeter_ClientStreamServer) error {

	for {
		req, err := streampb.Greeter_ClientStreamServer.Recv(res)
		if err != nil {
			return err
		}
		fmt.Println(req)
	}

}

// 双向流模式
func (s *Server) AllStreeam(allStream streampb.Greeter_AllStreeamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			req, err := streampb.Greeter_AllStreeamServer.Recv(allStream)
			if err != nil {
				return
			}
			fmt.Println(req)
		}
	}()

	go func() {
		defer wg.Done()
		i := 0
		for {
			i++
			_ = allStream.Send(&streampb.StreamResponse{
				Data: fmt.Sprintf("我是服务端发送的数据" ),
			})
			time.Sleep(time.Second)
			if i > 5 {
				break
			}
		}
	}()
	wg.Wait()
	return nil
}

func main() {

	g := grpc.NewServer()

	streampb.RegisterGreeterServer(g, &Server{})

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
	fmt.Println("启动成功")
}
