package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"lgo/pz-shop-server/user-srv/global"
	"lgo/pz-shop-server/user-srv/handler"
	"lgo/pz-shop-server/user-srv/initlalize"
	userpb "lgo/pz-shop-server/user-srv/proto"
	"lgo/pz-shop-server/user-srv/utils"

	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

var (
	IP        *string
	Port      *int
	serviceID = uuid.NewV4().String()
)

// 服务注册
func ServiceRegister(server *grpc.Server) *api.Client {
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
		ID:      serviceID,
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
	return client
}

func main() {
	initlalize.InitLogger()
	initlalize.InitConfig()
	initlalize.InitDB()

	server_info := global.ConfigYaml.ServerInfo

	IP = flag.String("ip", server_info.Host, "IP address")
	Port = flag.Int("port", 0, "Port number")

	flag.Parse()

	server := grpc.NewServer()

	userpb.RegisterUserServiceServer(server, &handler.UserServer{})

	if *Port == 0 {
		*Port, _ = utils.GetFreePort()
	}

	client := ServiceRegister(server)

	zap.S().Infof("user-src服务启动%s", fmt.Sprintf("%s:%d", *IP, *Port))

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))

	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	go func() {
		error := server.Serve(lis)
		if error != nil {
			panic("failed to serve:" + error.Error())
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
