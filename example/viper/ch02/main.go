package main

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

type ServerInfo struct {
	Host  string `mapstructure:"host"`
	Port  int    `mapstructure:"port"`
	Group string `mapstructure:"group"`
}

var _serverInfo = &ServerInfo{}
var configFileName string

func main() {
	//日志初始化
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	// 初步设计两个环境， 线上环境和开发环境，  开发环境就是测试环境
	//系统环境变量
	isDev := GetEnvInfo("PZSHOP_DEV")

	zap.S().Infof("在系统环境中读取，是否是开发环境: %v\n", isDev)
	configFilePrefix := "config"
	configFileName = fmt.Sprintf("%s-pro.yaml", configFilePrefix)

	if isDev {
		configFileName = fmt.Sprintf("%s-dev.yaml", configFilePrefix)
	}

	zap.S().Infof("当前的配置文件是: %s\n", configFileName)

	// 读取配置文件

	v := viper.New()

	readInConfigAndUnmarshal(v)

	// 动态监听 yaml的变化
	v.WatchConfig()

	// 监听配置变化
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置文件被修改: %s \n", e.Name)
		readInConfigAndUnmarshal(v)
	})

	time.Sleep(time.Second * 60)
}

// 读取配置文件
func readInConfigAndUnmarshal(v *viper.Viper) {
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Errorf("Fatal error config file: %s \n", err)
		panic(err)
	}

	if err := v.Unmarshal(_serverInfo); err != nil {
		zap.S().Errorf("Unmarshal config file error: %s \n", err)
		panic(err)
	}
	zap.S().Infof("读取到的配置信息是: %#v\n", _serverInfo)
}
