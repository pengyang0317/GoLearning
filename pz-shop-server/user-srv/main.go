package main

import (
	"flag"
	"fmt"
	"net"

	"lgo/pz-shop-server/user-srv/global"
	"lgo/pz-shop-server/user-srv/handler"
	"lgo/pz-shop-server/user-srv/initlalize"
	userpb "lgo/pz-shop-server/user-srv/proto"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func startServer() {
	server_info := global.ConfigYaml.ServerInfo

	IP := flag.String("ip", server_info.Host, "IP address")
	Port := flag.Int("port", server_info.Port, "Port number")

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

func main() {
	initlalize.InitLogger()
	initlalize.InitConfig()
	initlalize.InitDB()
	// startServer()

	server_info := global.ConfigYaml.ServerInfo

	IP := flag.String("ip", server_info.Host, "IP address")
	Port := flag.Int("port", server_info.Port, "Port number")

	flag.Parse()

	fmt.Printf("IP: %s, Port: %d", *IP, *Port)

	server := grpc.NewServer()

	userpb.RegisterUserServiceServer(server, &handler.UserServer{})

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	ConsulInfo := global.ConfigYaml.ConsulInfo
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
		ID:      "user-srv",
		Name:    "user-srv",
		Address: *IP,
		Port:    *Port,
		Tags:    []string{"pengze", "go", "user", "srv"},
		Check: &api.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%s:%d", *IP, *Port),
			GRPCUseTLS:                     false,
			Timeout:                        "5s",
			Interval:                       "5s",
			DeregisterCriticalServiceAfter: "15s",
		},
	}

	zap.S().Infof("我是注册GRPC服务:%s", fmt.Sprintf("%s:%d", *IP, *Port))
	err = client.Agent().ServiceRegister(registration)

	if err != nil {
		panic(err)
	}

	error := server.Serve(lis)

	if error != nil {
		panic("failed to serve:" + error.Error())
	}

}
