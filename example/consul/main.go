package main

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/consul/api"
)

const (
	IP   = "http://192.168.0.104"
	PORT = 8500
)

func Register(address, id, name string, port int, tag []string) error {
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", IP, PORT)

	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	// 生成检查对象
	check := &api.AgentServiceCheck{
		Interval:                       "5s",                              // 指定运行此检查的频率
		Timeout:                        "5s",                              // 超时时间
		HTTP:                           fmt.Sprintf("%s:8021/health", IP), // 健康检查HTTP请求
		Method:                         http.MethodGet,                    // 健康检查请求方法
		DeregisterCriticalServiceAfter: "30s",                             // 触发注销的时间
	}

	// 生成注册对象
	registration := api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Tags:    tag,
		Port:    port,
		Address: address,
		Check:   check,
	}

	// 注册服务
	err = client.Agent().ServiceRegister(&registration)
	return err
}

func main() {
	err := Register("192.168.0.104", "test-service", "test-service", 8081, []string{"番茄炒蛋", "test-service"})
	if err != nil {
		fmt.Println("服务注册失败")
	}
}
