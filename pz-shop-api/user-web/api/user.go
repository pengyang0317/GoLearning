package api

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"lgo/pz-shop-api/user-web/forms"
	"lgo/pz-shop-api/user-web/global"
	"lgo/pz-shop-api/user-web/middlewares"
	"lgo/pz-shop-api/user-web/models"
	userpb "lgo/pz-shop-api/user-web/proto"
	protoapipb "lgo/pz-shop-api/user-web/protoapi"

	"github.com/dgrijalva/jwt-go"
)

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	//将grpc的code转换成http的状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg:": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			case codes.AlreadyExists:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": e.Message(),
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Code(),
				})
			}
			return
		}
	}
}

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

	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户: %d", currentUser.ID)

	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)

	userList, err := global.UserServiceClient.GetUserList(ctx, &userpb.GetUserRequest{
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

	if respones, err := global.UserServiceClient.GetUserByMobile(context.Background(), &userpb.GetUserByMobileRequest{
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
		if rsp, err := global.UserServiceClient.CheckPassword(context.Background(), &userpb.CheckPasswordRequest{
			PassWord:          passwordLoginForm.PassWord,
			EncryptedPassWord: respones.User.PassWord,
		}); err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]string{
				"password": "密码错误",
			})
		} else {
			//step2: 生成token
			if rsp.Success {
				_createToken(ctx, respones.User)
			} else {
				ctx.JSON(http.StatusInternalServerError, map[string]string{
					"msg": "登录失败",
				})
			}
		}

	}

}

// 用户注册
func Register(c *gin.Context) {
	registerForm := protoapipb.RegisterForm{}

	if err := c.ShouldBind(&registerForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	val, err := global.Rdb.Get(registerForm.Mobile).Result()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"captcha": "验证码错误",
		})
		return
	}

	if val != registerForm.Code {
		c.JSON(http.StatusBadRequest, gin.H{
			"captcha": "验证码错误",
		})
		return
	}

	user, err := global.UserServiceClient.CreateUser(context.Background(), &userpb.CreateUserRequest{
		NickName: registerForm.Mobile,
		PassWord: registerForm.PassWord,
		Mobile:   registerForm.Mobile,
	})

	if err != nil {
		zap.S().Errorf("[Register] 查询 【新建用户失败】失败: %s", err.Error())
		HandleGrpcErrorToHttp(err, c)
		return
	}

	_createToken(c, user.User)
}

func GetUserDetail(c *gin.Context) {
	claims, _ := c.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户: %d", currentUser.ID)

	rsp, err := global.UserServiceClient.GetUserById(context.Background(), &userpb.GetUserByIdRequest{
		Id: int32(currentUser.ID),
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}

	user := rsp.User
	c.JSON(http.StatusOK, gin.H{
		"name":     user.NickName,
		"birthday": time.Unix(int64(user.BirthDay), 0).Format("2023-05-01"),
		"gender":   user.Gender,
		"mobile":   user.Mobile,
	})
}

func UpdateUser(ctx *gin.Context) {
	updateUserForm := forms.UpdateUserForm{}
	if err := ctx.ShouldBind(&updateUserForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户: %d", currentUser.ID)

	loc, _ := time.LoadLocation("Local")
	birthDay, _ := time.ParseInLocation("2023-05-01", updateUserForm.Birthday, loc)
	_, err := global.UserServiceClient.UpdateUser(context.Background(), &userpb.UpdateUserRequest{
		Id:       int32(currentUser.ID),
		NickName: updateUserForm.Name,
		Gender:   updateUserForm.Gender,
		BirthDay: uint64(birthDay.Unix()),
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

// 生成token
func _createToken(c *gin.Context, respones *userpb.User) {
	j := middlewares.NewJWT()
	claims := models.CustomClaims{
		ID:          uint(respones.Id),
		NickName:    respones.NickName,
		AuthorityId: uint(respones.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),               //签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
			Issuer:    "pengze",
		},
	}
	token, err := j.CreateToken(claims)

	if err != nil {
		zap.S().Errorf("获取token失败, %s", err.Error())
		c.JSON(http.StatusInternalServerError, map[string]string{
			"msg": "登录失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         respones.Id,
		"nick_name":  respones.NickName,
		"token":      token,
		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	})
}
