package config

type ConsulInfo struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type ServerInfo struct {
	Host string `mapstructure:"host" json:"host"`
}

type ServerConfig struct {
	ConsulInfo ConsulInfo `mapstructure:"consul_info" json:"consul_info"`
	ServerInfo ServerInfo `mapstructure:"server_info" json:"server_info"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	DataId    string `mapstructure:"dataId"`
	Group     string `mapstructure:"group"`
}
