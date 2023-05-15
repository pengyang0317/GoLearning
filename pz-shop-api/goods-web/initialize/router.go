package initialize

import (
	"lgo/pz-shop-api/goods-web/router"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	Router := gin.Default()

	ApiGroup := Router.Group("/g/v1")
	// router.InitGoodsRouter(ApiGroup)
	// router.InitCategoryRouter(ApiGroup)
	router.InitBannerRouter(ApiGroup)
	// router.InitBrandRouter(ApiGroup)

	return Router
}
