package initlalize

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"lgo/pz-shop-rpc/user-srv/config"
	"lgo/pz-shop-rpc/user-srv/global"
)

var (
	ServerConfig = global.ServerConfig
)

// 读取系统环境变量
func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

// 初始化读取配置文件
func initConfigFileName(v *viper.Viper) {
	isDev := GetEnvInfo("PZSHOP_DEV")

	configFilePrefix := "config"
	configFileName := fmt.Sprintf("%s-pro.yaml", configFilePrefix)

	if isDev {
		configFileName = fmt.Sprintf("%s-dev.yaml", configFilePrefix)
	}

	zap.S().Infof("当前的配置文件是: %s\n", configFileName)
	v.SetConfigFile(configFileName)
}

// 动态监听 配置文件的变化
func WatchConfig(v *viper.Viper, ServerInfo *config.ServerConfig) {

	// 动态监听 yaml的变化
	v.WatchConfig()

	// 监听配置变化
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("修改了配置文件: %s \n", e.Name)
		readInConfigAndUnmarshal(v, ServerInfo)
	})

}

// 读取配置文件并解析结构体
func readInConfigAndUnmarshal(v *viper.Viper, ServerInfo *config.ServerConfig) {
	if err := v.ReadInConfig(); err != nil {
		zap.S().Errorf("Fatal error config file: %s \n", err)
		panic(err)
	}

	if err := v.Unmarshal(ServerInfo); err != nil {
		zap.S().Errorf("Unmarshal config file error: %s \n", err)
		panic(err)
	}
}

func InitConfig() {

	v := viper.New()

	initConfigFileName(v)

	readInConfigAndUnmarshal(v, ServerConfig)

	WatchConfig(v, ServerConfig)

	zap.S().Infof("初始化配置文件成功: %#v \n", ServerConfig.ServerInfo)
}
