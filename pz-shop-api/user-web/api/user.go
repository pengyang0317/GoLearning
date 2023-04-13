package api

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"lgo/pz-shop-api/user-web/forms"
	"lgo/pz-shop-api/user-web/global"
	"lgo/pz-shop-api/user-web/middlewares"
	"lgo/pz-shop-api/user-web/models"
	userpb "lgo/pz-shop-api/user-web/proto"

	"github.com/dgrijalva/jwt-go"
)

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleValidatorError(ctx *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})

}

func GetUserList(ctx *gin.Context) {
	zap.S().Info("获取user列表")

	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		zap.S().Error("[GetUserList] 连接 [用户服务失败]", err)
	}

	userClient := userpb.NewUserServiceClient(conn)

	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户: %d", currentUser.ID)

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

func PassWordLogin(ctx *gin.Context) {
	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		zap.S().Error("[GetUserList] 连接 [用户服务失败]", err)
	}

	userClient := userpb.NewUserServiceClient(conn)

	zap.S().Info("密码登录")
	passwordLoginForm := forms.PassWordLoginForm{}
	if err := ctx.ShouldBind(&passwordLoginForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	if !store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, true) {
		zap.S().Info("验证码错误")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"captcha": "验证码错误",
		})
		return
	}

	zap.S().Info("登录逻辑")

	if respones, err := userClient.GetUserByMobile(context.Background(), &userpb.GetUserByMobileRequest{
		Mobile: passwordLoginForm.Mobile,
	}); err != nil {
		if e, ok := status.FromError(err); ok {
			zap.S().Errorf("获取用户信息失败, %s: code码:", e.Message(), e.Code())
			switch e.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusBadRequest, map[string]string{
					"mobile": "用户不存在",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]string{
					"mobile": "登录失败",
				})
			}
			return
		}
	} else {
		// step1: 校验密码
		if rsp, err := userClient.CheckPassword(context.Background(), &userpb.CheckPasswordRequest{
			PassWord:          passwordLoginForm.PassWord,
			EncryptedPassWord: respones.User.PassWord,
		}); err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]string{
				"password": "密码错误",
			})
		} else {
			//step2: 生成token
			if rsp.Success {
				j := middlewares.NewJWT()
				claims := models.CustomClaims{
					ID:          uint(respones.User.Id),
					NickName:    respones.User.NickName,
					AuthorityId: uint(respones.User.Role),
					StandardClaims: jwt.StandardClaims{
						NotBefore: time.Now().Unix(),               //签名的生效时间
						ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
						Issuer:    "pengze",
					},
				}
				token, err := j.CreateToken(claims)

				if err != nil {
					zap.S().Errorf("获取token失败, %s", err.Error())
					ctx.JSON(http.StatusInternalServerError, map[string]string{
						"msg": "登录失败",
					})
					return
				}

				ctx.JSON(http.StatusOK, gin.H{
					"id":         respones.User.Id,
					"nick_name":  respones.User.NickName,
					"token":      token,
					"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
				})

			} else {
				ctx.JSON(http.StatusInternalServerError, map[string]string{
					"msg": "登录失败",
				})
			}
		}

	}

}
