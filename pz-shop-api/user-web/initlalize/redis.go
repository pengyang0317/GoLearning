package initlalize

import (
	"fmt"
	"lgo/pz-shop-api/user-web/global"

	"github.com/go-redis/redis"
)

func InitRedis() {
	RedisInfo := global.ConfigYaml.RedisInfo
	global.Rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", RedisInfo.Host, RedisInfo.Port),
		DB:   RedisInfo.DB,
	})

}
