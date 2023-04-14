package main

import (
	"context"
	"fmt"
	hellopb "lgo/src/ch4/v2/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//grpc.WithTransportCredentials() 方法来设置传输层安全性，这里我们使用 insecure.NewCredentials() 方法来创建一个不安全的凭证，该凭证会告诉 gRPC 在传输层不使用任何安全机制。
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := hellopb.NewGreeterClient(conn)
	HelloResponse, err := c.SayHello(context.Background(), &hellopb.HelloRequest{Name: "world"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("HelloResponse: %v", HelloResponse)
}
