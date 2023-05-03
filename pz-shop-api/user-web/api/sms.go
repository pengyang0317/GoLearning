package api

import (
	"lgo/pz-shop-api/user-web/forms"
	"lgo/pz-shop-api/user-web/global"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SendSms(c *gin.Context) {

	sendSmsFormForm := forms.SendSmsForm{}
	RedisInfo := global.ConfigYaml.RedisInfo

	if err := c.ShouldBind(&sendSmsFormForm); err != nil {
		HandleValidatorError(c, err)
		return
	}

	err := global.Rdb.Set(sendSmsFormForm.Mobile, sendSmsFormForm.Code, time.Duration(RedisInfo.Expire)*time.Hour).Err()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "发送失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "发送成功",
	})
}
