package config

type MysqlInfo struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
}

type ServerInfo struct {
	Host  string `mapstructure:"host"`
	Port  int    `mapstructure:"port"`
	Group string `mapstructure:"group"`
	// JWTInfo JWTConfig `mapstructure:"jwt" json:"jwt"`
}

type ConsulInfo struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}

type ConfigYaml struct {
	MysqlInfo  MysqlInfo  `mapstructure:"mysql_info"`
	ServerInfo ServerInfo `mapstructure:"server_info"`
	ConsulInfo ConsulInfo `mapstructure:"consul_info"`
}
