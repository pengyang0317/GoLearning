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

func(GoodsServer) BannerList(ctx context.Context, req *emptypb.Empty) (*proto.BannerListResponse, error){
	var banners []model.Banner
	result := global.DB.Find(&banners)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "banner不存在")
	}
	var list []*proto.BannerResponse
	for _, banner := range banners {
		list = append(list, &proto.BannerResponse{
			Id:    banner.ID,
			Image: banner.Image,
			Url:   banner.Url,
			Index: banner.Index,
		})
	}
	return &proto.BannerListResponse{Data: list}, nil

}
func (s *GoodsServer) CreateBanner(ctx context.Context, req *proto.BannerRequest) (*proto.BannerResponse, error) {
	banner := model.Banner{}

	banner.Image = req.Image
	banner.Url = req.Url
	banner.Index = req.Index

	tx := global.DB.Create(&banner)
	if tx.Error != nil {
		zap.S().Errorf("global.DB.Create err: %v", tx.Error)
		return nil, tx.Error
	}
	return &proto.BannerResponse{Id: banner.ID}, nil
}

func (GoodsServer) DeleteBanner(ctx context.Context, req *proto.BannerRequest) (*emptypb.Empty, error) {
	tx := global.DB.Delete(&model.Banner{}, req.Id)
	if tx.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "删除失败")
	}
	return &emptypb.Empty{}, nil
}

func (GoodsServer) UpdateBanner(ctx context.Context, req *proto.BannerRequest) (*emptypb.Empty, error) {
	banner := &model.Banner{
		Image: req.Image,
		Url:   req.Url,
		Index: req.Index,
	}
	tx := global.DB.Model(banner).Where("id = ?", req.Id).Updates(banner)
	if tx.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "更新失败")
	}
	return &emptypb.Empty{}, nil
}
