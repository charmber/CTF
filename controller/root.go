package controller

import (
	"CTF/common"
	"CTF/model"
	"CTF/util"
	"github.com/gin-gonic/gin"
)

// CreatNotice 创建公告
func CreatNotice(c *gin.Context) {
	DB := common.GetDB()
	type Notice struct {
		Name    string `json:"name"`    //用户名
		Title   string `json:"title"`   //标题
		Author  string `json:"author"`  //作者
		Content string `json:"content"` //文章内容
		Time    string `json:"time"`    //时间
	}
	token := c.Query("token")
	_, ok := util.VerifyPermissions(token, c)
	if !ok {
		return
	}
	json := Notice{}
	c.BindJSON(&json)
	NewArt := model.Notice{
		Name:    json.Name,
		Title:   json.Title,
		Author:  json.Author,
		Content: json.Content,
		Time:    json.Time,
	}
	DB.Create(&NewArt)
	//返回结果
	c.JSON(200, gin.H{
		"code": 200,
		"data": NewArt,
		"msg":  "发布成功",
	})
}
