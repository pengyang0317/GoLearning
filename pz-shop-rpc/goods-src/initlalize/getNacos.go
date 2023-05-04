package initlalize

import (
	"encoding/json"
	"lgo/pz-shop-rpc/goods-src/global"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"
)

func InItNacos() {
	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   global.NacosConfig.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./tmp/nacos/log",
		CacheDir:            "./tmp/nacos/cache",
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})

	if err != nil {
		zap.S().Fatalf("初始化nacos配置中心失败: %v", err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group})

	if err != nil {
		zap.S().Fatalf("获取配置失败: %v", err)
	}

	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil {
		zap.S().Fatalf("读取nacos配置失败： %s", err.Error())
	}
	zap.S().Infof("获取配置成功: %v", global.ServerConfig)

	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group,
		OnChange: func(namespace, group, dataId, data string) {
			zap.S().Infof("group:%s, dataId:%s, data:%s", group, dataId, data)
			err = json.Unmarshal([]byte(content), &global.ServerConfig)
			if err != nil {
				zap.S().Fatalf("读取nacos配置失败： %s", err.Error())
			}
			zap.S().Infof("监听内的 -- 获取配置成功: %v", global.ServerConfig)
		},
	})
	if err != nil {
		zap.S().Fatalf("监听配置失败: %v", err)
	}

}
