package api

import (
	"net/http"
	"strconv"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	userpb "lgo/pz-shop-api/user-web/proto"
)

func GetUserList(ctx *gin.Context) {
	zap.S().Info("获取user列表")

	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		zap.S().Error("[GetUserList] 连接 [用户服务失败]", err)
	}

	userClient := userpb.NewUserServiceClient(conn)

	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)

	userList, err := userClient.GetUserList(ctx, &userpb.GetUserRequest{
		Size: int32(pSizeInt),
		Page: int32(pnInt),
	})

	if err != nil {
		zap.S().Error("[GetUserList] 获取 [用户列表失败]", err)
	}

	result := make([]interface{}, len(userList.Data))

	for i, value := range userList.Data {
		result[i] = userpb.User{
			Id:       value.Id,
			PassWord: value.PassWord,
			NickName: value.NickName,
			BirthDay: value.BirthDay,
			Role:     value.Role,
		}
	}

	zap.S().Info("获取user列表", userList)

	ctx.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "success",
		"data":  result,
		"total": userList.Totol,
	})

}
