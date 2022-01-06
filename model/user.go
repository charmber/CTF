package model

import "github.com/jinzhu/gorm"

type ComputerUser struct {
	gorm.Model
	UserName   string `gorm:"varchar(25);not null unique"` //用户名账号
	Email      string `gorm:"varchar(25);not null unique"` //邮箱
	NickName   string `gorm:"varchar(25);not null"`        //昵称
	Name       string `gorm:"varchar(25);not null"`        //姓名
	Number     string `gorm:"int(10);null unique"`         //学号
	Password   string `gorm:"varchar(255);not null"`       //密码
	Academy    string `gorm:"varchar(25);null"`            //学院
	Profession string `gorm:"varchar(25);null"`            //专业
	RatingAcm  int    `gorm:"int(8);null"`                 //acm评分
	RatingCtf  int    `gorm:"int(8);null"`                 //ctf评分
	BlogNum   int    `gorm:"int(8);null"`//博客数量
	BlogAcm    int    `gorm:"int(10);null"`//acm题解数量
	BlogCtf    int    `gorm:"int(8);null"`//ctf题解数量
	HeadImage  int    `gorm:"varchar(255);null"`//头像地址
	OpenNum   int    `gorm:"int(8);null"`//开源项目数量
	OpenStar   int    `gorm:"int(8);null"`//开源项目星星
	CreateTime string `gorm:"datetime not null"`
}
