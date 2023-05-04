package initlalize

import (
	"flag"
	"lgo/pz-shop-rpc/goods-src/global"
	"lgo/pz-shop-rpc/goods-src/utils"
)

func InitAddress() {
	IP := flag.String("ip", "0.0.0.0", "ip address")
	Port := flag.Int("port", 0, "port")

	if *IP == "0.0.0.0" {
		*IP = global.EnvInfo.IP
	}
	if *Port == 0 {
		*Port, _ = utils.GetFreePort()
	}

	global.StartServerIP = IP
	global.StartServerPort = Port
	flag.Parse()
}
