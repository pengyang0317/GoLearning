package main

import (
	"fmt"
	"time"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   8848,
		},
	}

	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "139058a1-5292-42ce-aed0-ac2c48a19896",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./tmp/nacos/log",
		CacheDir:            "./tmp/nacos/cache",
		LogLevel:            "debug",
	}
	// 创建动态配置客户端
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		fmt.Printf("动态配置客户端失败: %v\n", err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user-src-dev.json",
		Group:  "develop"})

	if err != nil {
		fmt.Printf("获取配置失败: %v\n", err)
	}

	fmt.Printf("获取配置成功: %v\n", content)

	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "user-src-dev.json",
		Group:  "develop",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
	if err != nil {
		fmt.Printf("监听配置失败: %v\n", err)
	}

	time.Sleep(300 * time.Second)
}
