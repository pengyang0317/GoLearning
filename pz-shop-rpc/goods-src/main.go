package main

import (
	"lgo/pz-shop-rpc/goods-src/initlalize"
)

func main() {
	IP, Port := initlalize.InitAddress()
	initlalize.InitLogger()

}
