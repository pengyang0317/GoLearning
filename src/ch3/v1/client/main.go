package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1.建立连接
	client, err := rpc.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Printf("dial error: %v", err)
	}
	var reply string

	// 2.调用远程方法
	err = client.Call("rpcServer.Hello", "world", &reply)

	if err != nil {
		fmt.Printf("call error: %v", err)
	}
	fmt.Println(reply)

}
