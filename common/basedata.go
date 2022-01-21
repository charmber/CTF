package common

import (
	"CTF/model"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// InitDB 数据库配置
func InitDB() *gorm.DB {
	driverName := "mysql"
	db, err := gorm.Open(driverName, "root:131420tj@(127.0.0.1:3306)/article?charset=utf8&parseTime=True&loc=Local&timeout=3600s")
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}

	//创建数据表
	//db.SetConnMaxLifetime(time.Duration(8*3600) * time.Second)
	db.AutoMigrate(&model.Article{}, &model.ComputerUser{}, &model.Problem{}, &model.Notice{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
