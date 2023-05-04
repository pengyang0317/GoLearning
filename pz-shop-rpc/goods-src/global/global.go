package global

import "lgo/pz-shop-rpc/goods-src/config"

type EnvInfoS struct {
	Env string
	IP  string
}

var (
	EnvInfo *EnvInfoS = &EnvInfoS{
		Env: "PZSHOP_DEV",
		IP:  "192.168.0.100",
	}
	ServerConfig    *config.ServerConfig = &config.ServerConfig{}
	StartServerIP   *string
	StartServerPort *int
	NacosConfig     *config.NacosConfig = &config.NacosConfig{}
)
