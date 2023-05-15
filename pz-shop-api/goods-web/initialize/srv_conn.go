package initialize

import (
	"fmt"
	"lgo/pz-shop-api/goods-web/global"
	"lgo/pz-shop-rpc/goods-src/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// InitSrvConn 初始化服务连接
func InitSrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.UserSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}

	global.GoodsSrvClient = proto.NewGoodsClient(userConn)
}
