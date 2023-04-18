package initlalize

import (
	"fmt"
	"lgo/pz-shop-api/user-web/global"
	userpb "lgo/pz-shop-api/user-web/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitSrvConn() {
	consulInfo := global.ConfigYaml.ConsulInfo
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, consulInfo.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}

	userSrvClient := userpb.NewUserServiceClient(userConn)
	global.UserServiceClient = userSrvClient
}
