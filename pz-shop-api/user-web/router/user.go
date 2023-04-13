package router

import (
	"lgo/pz-shop-api/user-web/api"

	"github.com/gin-gonic/gin"
	"lgo/pz-shop-api/user-web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("",middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("pwd_login", api.PassWordLogin)
	}
}
