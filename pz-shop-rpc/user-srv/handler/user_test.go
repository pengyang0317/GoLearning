package handler

import (
	"context"
	"fmt"
	"lgo/pz-shop-server/user-srv/global"
	userpb "lgo/pz-shop-server/user-srv/proto"
	"testing"
)

// 测试GetUserList
func TestGetUserList(t *testing.T) {
	//初始化数据库
	global.Init()
	//初始化handler
	userServer := new(UserServer)
	//调用GetUserList
	res, err := userServer.GetUserList(context.Background(), &userpb.GetUserRequest{
		Page: 2,
		Size: 2,
	})
	if err != nil {
		t.Errorf("GetUserList failed, err: %v", err)
	}
	fmt.Printf("res: %v", res)
}

// 测试GetUserByMobile

func TestGetUserByMobile(t *testing.T) {
	//初始化数据库
	global.Init()
	//初始化handler
	userServer := new(UserServer)

	//调用GetUserByMobile
	res, err := userServer.GetUserByMobile(context.Background(), &userpb.GetUserByMobileRequest{
		Mobile: "18211112222",
	})
	if err != nil {
		t.Errorf("GetUserByMobile failed, err: %v", err)
	}
	fmt.Printf("res: %v", res)
}

// 测试 GetUserById
func TestGetUserById(t *testing.T) {
	//初始化数据库
	global.Init()
	//初始化handler
	userServer := new(UserServer)

	//调用GetUserByMobile
	res, err := userServer.GetUserById(context.Background(), &userpb.GetUserByIdRequest{
		Id: 1,
	})

	if err != nil {
		t.Errorf("GetUserByMobile failed, err: %v", err)
	}

	fmt.Printf("res: %v", res)
}

// 测试 CreateUser
func TestCreateUser(t *testing.T) {
	//初始化数据库
	global.Init()
	//初始化handler
	userServer := new(UserServer)

	//调用CreateUser
	res, err := userServer.CreateUser(context.Background(), &userpb.CreateUserRequest{
		NickName: "test-NickName",
		Mobile:   "18210437835",
		PassWord: "123456",
	})

	if err != nil {
		t.Errorf("CreateUser failed, err: %v", err)
	}

	fmt.Printf("res: %v", res)
}
