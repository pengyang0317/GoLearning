package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func main() {
	debug := GetEnvInfo("PZSHOP_DEV")

	fmt.Printf("debug: %v", debug)

	configFilePrefix := "config"
	configFileName := fmt.Sprintf("viper/ch02/%s-pro.yaml", configFilePrefix)

	if debug {
		configFileName = fmt.Sprintf("viper/ch02/%s-debug.yaml", configFilePrefix)
	}
	fmt.Printf("configFileName: %s", configFileName)
}
