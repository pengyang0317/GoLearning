package user

import (
	"context"
	"fmt"
	userpb "lgo/pz-shop-server/user-srv/proto"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGetUserList(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	userClient := userpb.NewUserServiceClient(conn)
	rsp, err := userClient.GetUserList(context.Background(), &userpb.GetUserRequest{
		Page: 1,
		Size: 5,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)

}
