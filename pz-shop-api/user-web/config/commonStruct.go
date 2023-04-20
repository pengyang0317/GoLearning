package config

type ServerInfo struct {
	Host    string    `mapstructure:"host"`
	Port    int       `mapstructure:"port"`
	Group   string    `mapstructure:"group"`
	JWTInfo JWTConfig `mapstructure:"jwt" json:"jwt"`
	Name    string    `mapstructure:"name"`
}

type ConsulInfo struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}

type ConfigYaml struct {
	ServerInfo ServerInfo `mapstructure:"server_info"`
	ConsulInfo ConsulInfo `mapstructure:"consul_info"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}
