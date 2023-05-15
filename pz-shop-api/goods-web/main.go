package main

import (
	"fmt"
	"lgo/pz-shop-api/goods-web/global"
	"lgo/pz-shop-api/goods-web/initialize"
	"lgo/pz-shop-api/user-web/utils"
	"os"
	"os/signal"
	"syscall"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()

	Router := initialize.InitRouter()

	initialize.InitSrvConn()

	//随机端口号
	port, _ := utils.GetFreePort()
	global.ServerConfig.Port = port

	register_client := initialize.NewRegistryClient(
		global.ServerConfig.ConsulInfo.Host,
		global.ServerConfig.ConsulInfo.Port)

	serviceId := fmt.Sprintf("%s", uuid.NewV4())

	err := register_client.Register(
		global.ServerConfig.Host,
		global.ServerConfig.Port,
		global.ServerConfig.Name,
		global.ServerConfig.Tags,
		serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败:", err.Error())
	}
	zap.S().Debugf("启动服务器, 端口： %d", global.ServerConfig.Port)
	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
			zap.S().Panic("启动失败:", err.Error())
		}
	}()
	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = register_client.DeRegister(serviceId); err != nil {
		zap.S().Info("注销失败:", err.Error())
	} else {
		zap.S().Info("注销成功:")
	}

}
