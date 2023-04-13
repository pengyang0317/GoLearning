package config

type ServerInfo struct {
	Host    string    `mapstructure:"host"`
	Port    int       `mapstructure:"port"`
	Group   string    `mapstructure:"group"`
	JWTInfo JWTConfig `mapstructure:"jwt" json:"jwt"`
}

type ConfigYaml struct {
	ServerInfo ServerInfo `mapstructure:"server_info"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}
