package handler

import (
	"context"
	"lgo/pz-shop-rpc/goods-src/global"
	"lgo/pz-shop-rpc/goods-src/model"
	"lgo/pz-shop-rpc/goods-src/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GoodsServer) BrandList(ctx context.Context, req *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	var brands []model.Brand
	result := global.DB.Scopes(Paginate(int(req.PageNum), int(req.PageSize))).Find(&brands)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}
	var list []*proto.BrandInfoResponse
	for _, brand := range brands {
		list = append(list, &proto.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo, 
		})
	}
	return &proto.BrandListResponse{Data: list}, nil

}
func (s *GoodsServer) CreateBrand(ctx context.Context, req *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	if result := global.DB.Where("name = ?", req.Name).First(&model.Brand{}); result.RowsAffected > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌已存在")
	}
	brand := &model.Brand{
		Name: req.Name,
		Logo: req.Logo,
	}
	tx := global.DB.Create(brand)
	zap.S().Infof("tx: %v", tx)
	return &proto.BrandInfoResponse{Id: brand.ID}, nil
}

func (s *GoodsServer) UpdateBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	brand := &model.Brand{
		Name: req.Name,
		Logo: req.Logo,
	}
	tx := global.DB.Model(brand).Where("id = ?", req.Id).Updates(brand)
	if tx.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "更新失败")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) DeleteBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	tx := global.DB.Delete(&model.Brand{}, req.Id)
	if tx.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "删除失败")
	}
	return &emptypb.Empty{}, nil
}
