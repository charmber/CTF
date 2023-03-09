package model

import "github.com/jinzhu/gorm"

type SubmitProblem struct {
	gorm.Model
	Number    string `gorm:"int(10);null unique"` //账号
	ProblemId string `gorm:"int;null"`            //题目id
}
