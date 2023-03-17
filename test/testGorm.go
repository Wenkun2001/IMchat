package main

import (
	"IMchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:010226@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	db.AutoMigrate(&models.UserBasic{})

	//// Create
	//user := &models.UserBasic{}
	//user.Name = "lwk"
	//db.Create(user)
	//// Read
	//db.First(user, 1) // 根据整形主键查找
	//db.Model(user).Update("PassWord", 123456)
}
