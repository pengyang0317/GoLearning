package router

import (
	"lgo/pz-shop-api/goods-web/api/banners"

	"github.com/gin-gonic/gin"
)


func InitBannerRouter( r *gin.RouterGroup){
	BannerRouter := r.Group("banners")
		{
			BannerRouter.GET("", banners.List)          // 轮播图列表页
			// BannerRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), banners.Delete) // 删除轮播图
			// BannerRouter.POST("",  middlewares.JWTAuth(), middlewares.IsAdminAuth(), banners.New)       //新建轮播图
			// BannerRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), banners.Update) //修改轮播图信息
		}
}