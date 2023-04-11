package main

import (
	"context"
	"fmt"
	streampb "lgo/src/ch4/v3/proto"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 客户端流模式
func clientStream(c streampb.GreeterClient) {
	putS, _ := c.ClientStream(context.Background())
	i := 0
	for {
		i++
		_ = putS.Send(&streampb.StreamRequest{
			Data: fmt.Sprintf("我是客户端流数据%d", i),
		})
		time.Sleep(time.Second)
		if i > 5 {
			break
		}
	}
}

// 服务端流模式
func serverStream(c streampb.GreeterClient) {
	getS, _ := c.ServerStream(context.Background(), &streampb.StreamRequest{
		Data: "我是服务端流数据",
	})
	for {
		res, err := getS.Recv()
		if err != nil {
			break
		}
		fmt.Println(res)
	}
}

// 双向流模式
func twoWayStream(c streampb.GreeterClient, wg sync.WaitGroup) {

	allStream, _ := c.AllStreeam(context.Background())
	go func() {
		defer wg.Done()
		for {
			data, _ := allStream.Recv()
			fmt.Printf("%v\n", data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		i := 0
		for {
			i++
			if i > 5 {
				return
			}
			_ = allStream.Send(&streampb.StreamRequest{
				Data: fmt.Sprintf("发送数据"),
			})
			time.Sleep(time.Second)
		}

	}()
	wg.Wait()
}

func main() {

	// grpc.WithTransportCredentials() //方法来设置传输层安全性，这里我们使用 insecure.NewCredentials() 方法来创建一个不安全的凭证，该凭证会告诉 gRPC 在传输层不使用任何安全机制。
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := streampb.NewGreeterClient(conn)

	// clientStream(c)
	// serverStream(c)
	wg := sync.WaitGroup{}
	wg.Add(2)
	twoWayStream(c, wg)
	time.Sleep(8 * time.Second)

}
