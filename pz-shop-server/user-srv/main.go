package main

import (
	"flag"
	"fmt"
	"net"

	"lgo/pz-shop-server/user-srv/global"
	"lgo/pz-shop-server/user-srv/handler"
	userpb "lgo/pz-shop-server/user-srv/proto"

	"google.golang.org/grpc"
)

func main() {

	//初始化数据库
	global.Init()

	IP := flag.String("ip", "0.0.0.0", "IP address")
	Port := flag.Int("port", 8000, "Port number")

	flag.Parse()

	fmt.Printf("IP: %s, Port: %d", *IP, *Port)

	server := grpc.NewServer()

	userpb.RegisterUserServiceServer(server, &handler.UserServer{})

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	error := server.Serve(lis)

	if error != nil {
		panic("failed to serve:" + error.Error())
	}
}
