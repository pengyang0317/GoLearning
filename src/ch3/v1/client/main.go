package main

import (
	"fmt"
	"net/rpc"
)

const PORT = ":9000"

func main() {
	// 1.连接 RPC 服务端：使用 rpc.Dial 方法连接指定的 RPC 服务端，传入 tcp 协议和服务端的端口号（PORT）。
	client, err := rpc.Dial("tcp", PORT)
	if err != nil {
		fmt.Printf("dial error: %v", err)
	}
	var reply string

	// 2.调用远程方法：使用 client.Call 方法调用远程方法，传入方法名、参数和返回值指针。这里的方法名为 "rpcServer.Hello"，表示调用名为 "Hello" 的方法，该方法属于名为 "rpcServer" 的 RPC 服务。参数为 "pengze"，表示将 "pengze" 作为参数传递给远程方法。返回值为 reply，是一个字符串类型的指针，用于接收远程方法的返回结果。
	err = client.Call("rpcServer.Hello", "pengze", &reply)

	if err != nil {
		fmt.Printf("call error: %v", err)
	}
	fmt.Println(reply)

}
