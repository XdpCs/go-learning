package main

import (
	"GORM/config"
	"GORM/orm"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := gorm.Open(mysql.Open(config.DBConfig.URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	person := orm.Person{Username: "xdpcs", Email: "451679250@qq.com"}
	// 用指定的字段创建记录
	result := db.Table("person").Select("username", "email").Create(&person)
	fmt.Println(result.RowsAffected)
	result = db.Table("person").Omit("username", "user_id").Create(&person)
	fmt.Println(result.RowsAffected)
}
