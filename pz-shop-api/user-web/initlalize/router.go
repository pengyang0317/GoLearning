package initlalize

import (
	"lgo/pz-shop-api/user-web/middlewares"
	"lgo/pz-shop-api/user-web/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	//配置跨域
	Router.Use(middlewares.Cors())

	ApiGroup := Router.Group("/pengze/v1")

	router.InitUserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)

	return Router
}
