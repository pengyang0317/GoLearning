package initlalize

import "flag"

func InitAddress() (*string, *int) {
	IP := flag.String("ip", "0.0.0.0", "ip address")
	Port := flag.Int("port", 0, "port")
	flag.Parse()
	return IP, Port
}
