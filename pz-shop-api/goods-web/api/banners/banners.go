package banners

import (
	"context"
	"lgo/pz-shop-api/goods-web/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
)

func List(ctx *gin.Context) {

	res, err := global.GoodsSrvClient.BannerList(context.Background(), &empty.Empty{})

	zap.S().Info(res)
	zap.S().Info(err)

	// result := make([]interface{}, 0)
	// for _, value := range rsp.Data {
	// 	reMap := make(map[string]interface{})
	// 	reMap["id"] = value.Id
	// 	reMap["index"] = value.Index
	// 	reMap["image"] = value.Image
	// 	reMap["url"] = value.Url

	// 	result = append(result, reMap)
	// }

	// ctx.JSON(http.StatusOK, result)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
