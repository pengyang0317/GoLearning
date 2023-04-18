package main

import (
	"fmt"
	"lgo/pz-shop-api/user-web/initlalize"

	"go.uber.org/zap"

	"lgo/pz-shop-api/user-web/global"
)

func main() {
	// 初始化Logger
	initlalize.InitLogger()

	// 初始化配置文件
	initlalize.InitConfig()

	// 初始化翻译器
	if err := initlalize.InitTrans("zh"); err != nil {
		zap.S().Errorf("初始化翻译器失败，err:%s", err.Error())
		panic(err)
	}
	// 初始化路由
	Router := initlalize.Routers()

	// 初始化注册器
	if err := initlalize.InitRegister(); err != nil {
		zap.S().Errorf("初始化注册器失败，err:%s", err.Error())
		panic(err)
	}

	// 注册服务
	initlalize.InitSrvConn()

	zap.S().Debugf("启动服务器, 端口： %d", global.ConfigYaml.ServerInfo.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ConfigYaml.ServerInfo.Port)); err != nil {
		zap.S().Errorf("启动服务失败，err:%s", err.Error())
	}
}
