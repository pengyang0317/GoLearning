package handler

import (
	"context"
	"encoding/json"
	"lgo/pz-shop-rpc/goods-src/global"
	"lgo/pz-shop-rpc/goods-src/model"
	"lgo/pz-shop-rpc/goods-src/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (GoodsServer) GetAllCategorysList(ctx context.Context, req *emptypb.Empty) (*proto.CategoryListResponse, error) {
	var categorys []model.Category
	global.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categorys)
	b, _ := json.Marshal(&categorys)
	return &proto.CategoryListResponse{JsonData: string(b)}, nil
}

func (s *GoodsServer) CreateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	// category := model.Category{}
	cMap := map[string]interface{}{}
	cMap["name"] = req.Name
	cMap["is_tab"] = req.IsTab
	cMap["level"] = req.Level

	if req.Level != 1 {
		cMap["parent_category_id"] = req.ParentCategory
	}

	tx := global.DB.Model(&model.Category{}).Create(cMap)
	if tx.Error != nil {
		zap.S().Errorf("create category error: %v", tx.Error)
		return nil, tx.Error
	}

	return &proto.CategoryInfoResponse{}, nil
}

func (GoodsServer) DeleteCategory(ctx context.Context, req *proto.DeleteCategoryRequest) (*emptypb.Empty, error) {
	if tx := global.DB.Delete(&model.Category{}, req.Id); tx.RowsAffected == 0 {
		zap.S().Errorf("delete category error: %v", tx.Error)
		return nil, status.Errorf(codes.NotFound, "删除失败")
	}
	return &emptypb.Empty{}, nil
}

func (GoodsServer) UpdateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*emptypb.Empty, error) {
	global.DB.Model(&model.Category{}).Where("id = ?", req.Id).Updates(
		map[string]interface{}{
			"name":               req.Name,
			"is_tab":             req.IsTab,
			"level":              req.Level,
			"parent_category_id": req.ParentCategory,
		},
	)

	return &emptypb.Empty{}, nil
}
