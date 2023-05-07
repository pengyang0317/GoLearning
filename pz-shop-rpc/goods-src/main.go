package main

import (
	"fmt"
	"lgo/pz-shop-rpc/goods-src/global"
	"lgo/pz-shop-rpc/goods-src/handler"
	"lgo/pz-shop-rpc/goods-src/initlalize"
	"lgo/pz-shop-rpc/goods-src/proto"
	"net"
	"os"
	"os/signal"
	"syscall"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	serviceID = uuid.NewV4().String()
)

func main() {
	initlalize.InitAddress()
	initlalize.InitLogger()

	initlalize.InitConfig()
	initlalize.InItNacos()

	initlalize.InitDB()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *global.StartServerIP, *global.StartServerPort))

	if err != nil {
		zap.S().Errorf("net.Listen err: %v", err)
		panic(err)
	}
	server := grpc.NewServer()
	proto.RegisterGoodsServer(server, &handler.GoodsServer{})

	client := initlalize.InitConsul(server, serviceID)

	if err != nil {
		zap.S().Errorf("api.NewClient err: %v", err)
		panic(err)
	}

	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = client.Agent().ServiceDeregister(serviceID); err != nil {
		zap.S().Info("注销失败")
	}
	zap.S().Info("注销成功")
}
