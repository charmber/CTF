package model

import "github.com/jinzhu/gorm"

type Problem struct {
	gorm.Model
	ProblemName string `gorm:"varchar(255);not null"`
	ProblemType string `gorm:"varchar(25);not null"`
	Answer string `gorm:"varchar(255);not null"`
	ProblemPassNumber string `gorm:"varchar(25);not null"`
	ProblemSubmitNUmber string `gorm:"varchar(255);not null"`
	ProblemABloodId string `gorm:"varchar(25);not null"`
	Display string `gorm:"int(8);not null;unique"`
	FlagBool string `gorm:"varchar(25)"`
	Tips string `gorm:"varchar(255);"`
}
