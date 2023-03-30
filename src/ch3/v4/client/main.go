package main

import (
	"fmt"
	cproxy "lgo/src/ch3/v4/c_proxy"
)

func main() {
	// 1.建立连接
	client, err := cproxy.NewRpcServer("localhost:9000")
	if err != nil {
		fmt.Printf("dial error: %v", err)
	}

	var reply string

	// 2.调用远程方法
	err = client.Hello("pengze", &reply)
	if err != nil {
		fmt.Printf("call error: %v", err)
	}

	fmt.Println(reply)

}
