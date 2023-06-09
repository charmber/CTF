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

func RootLogin(c *gin.Context) {
	DB := common.GetDB()
	//获取参数
	/*number:=lg.PostForm("number")
	password:=lg.PostForm("password")*/
	//tp:1账号，2用户名，3邮箱
	type ComputerUser struct {
		Tp       string `json:"tp"`
		Number   string `json:"number"`
		UserName string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var user ComputerUser
	json := ComputerUser{}
	c.BindJSON(&json)
	if json.Tp == "1" {
		//数据验证
		if len(json.Password) <= 6 {
			c.JSON(200, gin.H{"code": 403, "msg": "密码不能少于6位"})
			return
		}
		//判断是否存在
		result := DB.Where("number=? AND permissions==1", json.Number).Find(&user)
		if result.RowsAffected == 0 {
			c.JSON(200, gin.H{"code": 403, "msg": "用户不存在"})
			return
		}
	} else if json.Tp == "2" {
		//数据验证
		if len(json.Password) <= 6 {
			c.JSON(200, gin.H{"code": 403, "msg": "密码不能少于6位"})
			return
		}
		//判断是否存在
		result := DB.Where("user_name=? AND permissions==1", json.UserName).Find(&user)
		if result.RowsAffected == 0 {
			c.JSON(200, gin.H{"code": 403, "msg": "用户不存在"})
			return
		}
	} else if json.Tp == "3" {
		//数据验证
		if len(json.Password) <= 6 {
			c.JSON(200, gin.H{"code": 403, "msg": "密码不能少于6位"})
			return
		}
		//判断是否存在
		result := DB.Where("email=? AND permissions==1", json.Email).Find(&user)
		if result.RowsAffected == 0 {
			c.JSON(200, gin.H{"code": 403, "msg": "用户不存在"})
			return
		}
	} else {
		c.JSON(200, gin.H{"code": 404, "msg": "不存在登录方式"})
	}

	if Check(json.Password, user.Password) != true {
		c.JSON(200, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	//if json.Password!=user.Password{
	//	lg.JSON(http.StatusBadRequest,gin.H{"code":400,"msg":"密码错误"})
	//		return
	//}

	//发放token
	token, _ := util.CreateToken(json.Number, 1)

	//返回结果
	c.JSON(200, gin.H{
		"code":     200,
		"token":    token,
		"email":    user.Email,
		"number":   user.Number,
		"username": user.UserName,
		"msg":      "登录成功",
	})
}
