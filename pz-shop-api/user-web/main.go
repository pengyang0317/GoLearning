package main

import (
	"lgo/pz-shop-api/user-web/initlalize"

	"go.uber.org/zap"
)

func main() {
	var PROT = ":8900"
	initlalize.InitLogger()

	initlalize.InitConfig()

	Router := initlalize.Routers()

	zap.S().Debugf("启动服务，端口：%s", PROT)
	if err := Router.Run(PROT); err != nil {
		zap.S().Errorf("启动服务失败，err:%s", err.Error())
	}
}
