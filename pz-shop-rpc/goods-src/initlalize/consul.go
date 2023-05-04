package initlalize

import (
	"fmt"
	"lgo/pz-shop-rpc/goods-src/global"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func ServiceRegister(server *grpc.Server, serviceID string) *api.Client {
	ConsulInfo := global.ServerConfig.ConsulInfo
	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	//服务注册
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", ConsulInfo.Host,
		ConsulInfo.Port)

	client, err := api.NewClient(cfg)

	if err != nil {
		panic(err)
	}

	//生成注册对象
	registration := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "goods-srv",
		Address: *global.StartServerIP,
		Port:    *global.StartServerPort,
		Tags:    []string{"pengze", "go", "user", "srv"},
		Check: &api.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%s:%d", *global.StartServerIP, *global.StartServerPort),
			GRPCUseTLS:                     false,
			Timeout:                        "5s",
			Interval:                       "5s",
			DeregisterCriticalServiceAfter: "15s",
		},
	}

	zap.S().Infof("我是注册GRPC服务:%s", fmt.Sprintf("%s:%d", *global.StartServerIP, *global.StartServerPort))
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return client
}

func InitConsul(server *grpc.Server, serviceID string) *api.Client {
	return ServiceRegister(server, serviceID)
}
