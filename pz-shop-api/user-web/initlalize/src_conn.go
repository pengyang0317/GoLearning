package initlalize

import (
	"fmt"
	"lgo/pz-shop-api/user-web/global"
	userpb "lgo/pz-shop-api/user-web/proto"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// func InitSrvConn(){
// 	consulInfo := global.ServerConfig.ConsulInfo
// 	userConn, err := grpc.Dial(
// 		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.UserSrvInfo.Name),
// 		grpc.WithInsecure(),
// 		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
// 	)
// 	if err != nil {
// 		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
// 	}

// 	userSrvClient := proto.NewUserClient(userConn)
// 	global.UserSrvClient = userSrvClient
// }

func InitSrvConn() {
	// conn, err := grpc.Dial("192.168.122.114:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	zap.S().Error("[GetUserList] 连接 [用户服务失败]", err)
	// }

	// userSrvClient := userpb.NewUserServiceClient(conn)
	// global.UserServiceClient = userSrvClient
	consulInfo := global.ConfigYaml.ConsulInfo

	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ConfigYaml.ServerInfo.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}

	userSrvClient := userpb.NewUserServiceClient(userConn)
	global.UserServiceClient = userSrvClient
}
