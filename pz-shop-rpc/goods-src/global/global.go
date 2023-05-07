package global

import (
	"lgo/pz-shop-rpc/goods-src/config"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type EnvInfoS struct {
	Env string
	IP  string
}

var (
	DB      *gorm.DB
	EnvInfo *EnvInfoS = &EnvInfoS{
		Env: "PZSHOP_DEV",
		IP:  "192.168.0.100",
	}
	ServerConfig    *config.ServerConfig = &config.ServerConfig{}
	StartServerIP   *string
	StartServerPort *int
	NacosConfig     *config.NacosConfig = &config.NacosConfig{}
)

func InitCeshiDB() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/pzshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}

	// DB.AutoMigrate(&model.Brand{})
}
