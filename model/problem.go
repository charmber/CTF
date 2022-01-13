package model

import "github.com/jinzhu/gorm"

type Problem struct {
	gorm.Model
	ProblemName         string `gorm:"varchar(255);not null"`
	ProblemType         string `gorm:"varchar(25);not null"`
	Answer              string `gorm:"varchar(255);not null"`
	ProblemPassNumber   string `gorm:"varchar(25);not null"`  //问题通过次数
	ProblemSubmitNumber string `gorm:"varchar(255);not null"` //问题提交次数
	ProblemABloodId     string `gorm:"varchar(25);not null"`
	Display             string `gorm:"int(8);not null;unique"` //问题编号
	FlagBool            string `gorm:"varchar(25)"`            //动态flag开关
	Tips                string `gorm:"varchar(255);"`          //问题提示
}
