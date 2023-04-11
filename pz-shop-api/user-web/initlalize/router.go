package initlalize

import (
	"lgo/pz-shop-api/user-web/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	ApiGroup := Router.Group("/pengze/v1")

	router.InitUserRouter(ApiGroup)

	return Router
}
