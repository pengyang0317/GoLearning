package global

import (
	"lgo/pz-shop-api/user-web/config"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis"

	userpb "lgo/pz-shop-api/user-web/proto"
)

var (
	Trans      ut.Translator
	ConfigYaml *config.ConfigYaml = &config.ConfigYaml{}

	UserServiceClient userpb.UserServiceClient

	Rdb  		   *redis.Client
)
