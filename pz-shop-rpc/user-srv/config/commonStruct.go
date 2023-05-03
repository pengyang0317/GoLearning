package config

type MysqlInfo struct {
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"name" json:"name"`
}

type ServerInfo struct {
	Host  string `mapstructure:"host" json:"host"`
	Port  int    `mapstructure:"port" json:"port"`
	Group string `mapstructure:"group" json:"group"`
	// JWTInfo JWTConfig `mapstructure:"jwt" json:"jwt"`
}

type ConsulInfo struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type ServerConfig struct {
	MysqlInfo  MysqlInfo  `mapstructure:"mysql_info" json:"mysql_info"`
	ServerInfo ServerInfo `mapstructure:"server_info" json:"server_info"`
	ConsulInfo ConsulInfo `mapstructure:"consul_info" json:"consul_info"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	DataId    string `mapstructure:"dataId"`
	Group     string `mapstructure:"group"`
}
