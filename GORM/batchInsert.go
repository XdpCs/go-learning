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
		// 使用CreateBatchSize 选项初始化 GORM 时，所有的创建& 关联 INSERT 都将遵循该选项
		// CreateBatchSize: 100,
	})
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	var persons = []orm.Person{{Username: "xdp1"}, {Username: "xdp2"}, {Username: "xdp3"}}
	// 批量全部插入
	db.Table("person").Create(&persons)
	// 该批量插入会分成两个语句进行插入
	db.Table("person").CreateInBatches(&persons, 2)
	for _, person := range persons {
		fmt.Println(person.ID)
	}
}
