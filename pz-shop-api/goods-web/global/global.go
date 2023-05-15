package global

import "lgo/pz-shop-rpc/goods-src/proto"

type NacosConfigC struct {
	Host      string `mapstructure:"host" json:"host" yaml:"host"`
	Port      uint64 `mapstructure:"port" json:"port" yaml:"port"`
	Namespace string `mapstructure:"namespace" json:"namespace" yaml:"namespace"`
	User      string `mapstructure:"user" json:"user" yaml:"user"`
	Password  string `mapstructure:"password" json:"password" yaml:"password"`
	DataId    string `mapstructure:"dataid" json:"dataid" yaml:"dataid"`
	Group     string `mapstructure:"group" json:"group" yaml:"group"`
}

type ServerConfigC struct {
	ConsulInfo  ConsulInfo  `mapstructure:"consul" json:"consul" yaml:"consul"`
	UserSrvInfo UserSrvInfo `mapstructure:"user_srv" json:"user_srv" yaml:"user_srv"`
	Host        string      `mapstructure:"host" json:"host" yaml:"host"`
	Port        int         `mapstructure:"port" json:"port" yaml:"port"`
	Name        string      `mapstructure:"name" json:"name" yaml:"name"`
	Tags        []string    `mapstructure:"tags" json:"tags" yaml:"tags"`
}

type ConsulInfo struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
	Name string `mapstructure:"name" json:"name" yaml:"name"`
	Tags []string `mapstructure:"tags" json:"tags" yaml:"tags"`
}
type UserSrvInfo struct {
	Name string `mapstructure:"name" json:"name" yaml:"name"`
}

var (
	NacosConfig    *NacosConfigC  = &NacosConfigC{}
	ServerConfig   *ServerConfigC = &ServerConfigC{}
	GoodsSrvClient proto.GoodsClient
)
