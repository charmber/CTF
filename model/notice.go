package model

import "github.com/jinzhu/gorm"

type Notice struct {
	gorm.Model
	Name      string `gorm:"varchar(25);not null"`   //用户名
	Title     string `gorm:"varchar(255);not null"`  //标题
	Author    string `gorm:"varchar(25);not null"`   //作者
	Content   string `gorm:"text;not null"`          //文章内容
	Time      string `gorm:"varchar(25);not null"`   //时间
	Image     string `gorm:"varchar(1024);not null"` //图片地址
	Recommend int    `gorm:"int;default:0"`          //浏览量
}
