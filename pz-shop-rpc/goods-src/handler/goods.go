package handler

import (
	"lgo/pz-shop-rpc/goods-src/proto"
)

type GoodsServer struct {
	proto.UnimplementedGoodsServer
}

// // 品牌分类
// CategoryBrandList(context.Context, *CategoryBrandFilterRequest) (*CategoryBrandListResponse, error)
// // 通过category获取brands
// GetCategoryBrandList(context.Context, *CategoryInfoRequest) (*BrandListResponse, error)
// CreateCategoryBrand(context.Context, *CategoryBrandRequest) (*CategoryBrandResponse, error)
// DeleteCategoryBrand(context.Context, *CategoryBrandRequest) (*emptypb.Empty, error)
// UpdateCategoryBrand(context.Context, *CategoryBrandRequest) (*emptypb.Empty, error)
