package test

import (
	"CTF/common"
	"github.com/gin-gonic/gin"
)

func Test(f *gin.Context) {
	DB := common.GetDB()
	type Article struct {
		Name  string
		Title string
	}
	var cat []Article
	DB.Limit(5).Order("created_at desc").Select([]string{"name", "title"}).Find(&cat)
	f.JSON(200, gin.H{
		"data": cat,
	})
}
