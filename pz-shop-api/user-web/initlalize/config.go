package initlalize

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	debug := GetEnvInfo("PZSHOP_DEV")

	configFilePrefix := "config"
	configFileName := fmt.Sprintf("viper/ch02/%s-pro.yaml", configFilePrefix)

	if debug {
		configFileName = fmt.Sprintf("viper/ch02/%s-debug.yaml", configFilePrefix)
	}

	zap.S().Infof("configFileName: %s", configFileName)
}
