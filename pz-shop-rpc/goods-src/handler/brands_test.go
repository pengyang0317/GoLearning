package handler

import (
	"context"
	"lgo/pz-shop-rpc/goods-src/global"
	"lgo/pz-shop-rpc/goods-src/initlalize"
	"lgo/pz-shop-rpc/goods-src/proto"
	"testing"

	"go.uber.org/zap"
)

// 测试 CreateBrand
func TestCreateBrand(t *testing.T) {
	global.InitCeshiDB()

	GoodsServer := new(GoodsServer)
	res, err := GoodsServer.CreateBrand(context.Background(), &proto.BrandRequest{
		Name: "test3",
		Logo: "test",
	})

	if err != nil {
		zap.S().Errorf("CreateBrand failed, err: %v", err)
	}
	zap.S().Infof("res: %v", res)
}

// 测试 UpdateBrand
func TestUpdateBrand(t *testing.T) {
	global.InitCeshiDB()
	initlalize.InitLogger()
	GoodsServer := new(GoodsServer)
	res, err := GoodsServer.UpdateBrand(context.Background(), &proto.BrandRequest{
		Id:   4,
		Name: "test3",
		Logo: "abcd",
	})

	if err != nil {
		zap.S().Errorf("UpdateBrand failed, err: %v", err)
	}
	zap.S().Infof("res: %v", res)
}

// 测试 DeleteBrand
func TestDeleteBrand(t *testing.T) {
	global.InitCeshiDB()
	initlalize.InitLogger()
	GoodsServer := new(GoodsServer)
	res, err := GoodsServer.DeleteBrand(context.Background(), &proto.BrandRequest{
		Id: 4,
	})
	if err != nil {
		zap.S().Errorf("DeleteBrand failed, err: %v", err)
	}
	zap.S().Infof("res: %v", res)
}

// 测试 BrandList
func TestBrandList(t *testing.T) {
	global.InitCeshiDB()
	initlalize.InitLogger()
	GoodsServer := new(GoodsServer)
	res, err := GoodsServer.BrandList(context.Background(), &proto.BrandFilterRequest{
		PageNum:  1,
		PageSize: 2,
	})
	if err != nil {
		zap.S().Errorf("BrandList failed, err: %v", err)
	}
	zap.S().Infof("res: %v", res)
}
