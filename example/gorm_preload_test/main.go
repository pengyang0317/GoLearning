package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// `源自官网示例`
type User struct {
	Id       int     `gorm:"primaryKey"`
	Username string  `gorm:"column:username"`
	Orders   []Order `gorm:"foreignKey:UserID"` // 一对多, 一个用户有多个订单
}

type Order struct {
	Id     int     `gorm:"primaryKey"`
	UserID uint    `gorm:"column:user_id"`
	Price  float64 `gorm:"column:price"`
}

func main() {
	// 数据库连接
	dsn := "root:123456@(127.0.0.1:3306)/example?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // 禁用彩色打印
			},
		),
	})
	if err != nil {
		panic(err)
	}

	var a []User
	db.Preload("Orders").Find(&a)
	fmt.Println(a)

	// 初始化表
	// _ = db.AutoMigrate(User{}, Order{})

	// 创建数据
	// _ = CreatUser(db)
	// _ = CreatOrder(db)
}

// func CreatUser(db *gorm.DB) (err error) {
// 	err = db.Create(&[]User{
// 		{Username: "little_A"},
// 		{Username: "little_B"},
// 		{Username: "little_C"},
// 		{Username: "little_D"},
// 	}).Error
// 	return
// }

// func CreatOrder(db *gorm.DB) (err error) {
// 	err = db.Create(&[]Order{
// 		{UserID: 1, Price: 1},
// 		{UserID: 1, Price: 2},
// 		{UserID: 1, Price: 3},
// 		{UserID: 1, Price: 4},
// 		{UserID: 2, Price: 5},
// 		{UserID: 2, Price: 6},
// 		{UserID: 2, Price: 7},
// 		{UserID: 3, Price: 8},
// 		{UserID: 3, Price: 9},
// 		{UserID: 4, Price: 10},
// 	}).Error
// 	return
// }
