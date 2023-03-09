package controller

import (
	"CTF/common"
	"CTF/model"
	"CTF/util"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"strings"
)

// Register 注册用户
func Register(c *gin.Context) {

	DB := common.GetDB()
	//获取参数
	/*number:=lg.PostForm("number")
	password:=lg.PostForm("password")*/
	type Username struct {
		UserName   string `json:"user_name"`  //用户名账号
		Email      string `json:"email"`      //邮箱
		NickName   string `json:"nick_name"`  //昵称
		Name       string `json:"name"`       //姓名
		Number     string `json:"number"`     //学号
		Password   string `json:"password"`   //密码
		Academy    string `json:"academy"`    //学院
		Profession string `json:"profession"` //专业
		CreateTime string `json:"create_time"`
	}
	json := Username{}
	c.BindJSON(&json)

	//数据验证
	if len(json.UserName) == 0 {
		c.JSON(422, gin.H{
			"code": 422,
			"msg":  "用户名账号必须填写(英文/数字/符号)",
		})
	}
	if len(json.Email) == 0 {
		c.JSON(422, gin.H{
			"code": 422,
			"msg":  "邮箱账号必须填写",
		})
	}
	if len(json.Name) == 0 {
		c.JSON(422, gin.H{
			"code": 422,
			"msg":  "用户名账号必须填写(英文/数字/符号)",
		})
	}
	if len(json.Password) <= 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	//如果人名为空，就随机一个10位字符串
	if len(json.NickName) == 0 {
		json.Name = util.Random(10)
	}

	//验证是否用户存在
	if InUserName(DB, json.UserName) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户名已经存在"})
		return
	}
	if InEmail(DB, json.Email) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "邮箱用户已经存在"})
		return
	}
	if len(json.Number) != 0 {
		if Innumber(DB, json.Number) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "学号已经存在"})
			return
		}
	}

	//创建用户
	//加密密码
	md := Encode(json.Password)
	newUser := model.ComputerUser{
		Name:       json.Name,
		Number:     json.Number,
		Password:   md,
		UserName:   json.UserName,
		Email:      json.Email,
		NickName:   json.NickName,
		Academy:    json.Academy,
		Profession: json.Profession,
		CreateTime: json.CreateTime,
	}
	DB.Create(&newUser)

	//返回结果
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

// Login 用户登录
func Login(lg *gin.Context) {

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
	lg.BindJSON(&json)
	if json.Tp == "1" {
		//数据验证
		if len(json.Password) <= 6 {
			lg.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
			return
		}
		//判断是否存在
		result := DB.Where("number=?", json.Number).Find(&user)
		if result.RowsAffected == 0 {
			lg.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
			return
		}
	} else if json.Tp == "2" {
		//数据验证
		if len(json.Password) <= 6 {
			lg.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
			return
		}
		//判断是否存在
		result := DB.Where("user_name=?", json.UserName).Find(&user)
		if result.RowsAffected == 0 {
			lg.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
			return
		}
	} else if json.Tp == "3" {
		//数据验证
		if len(json.Password) <= 6 {
			lg.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
			return
		}
		//判断是否存在
		result := DB.Where("email=?", json.Email).Find(&user)
		if result.RowsAffected == 0 {
			lg.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
			return
		}
	} else {
		lg.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "不存在登录方式"})
	}

	if Check(json.Password, user.Password) != true {
		lg.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	//if json.Password!=user.Password{
	//	lg.JSON(http.StatusBadRequest,gin.H{"code":400,"msg":"密码错误"})
	//		return
	//}

	//发放token
	token, _ := util.CreateToken(json.Number, 0)

	//返回结果
	lg.JSON(200, gin.H{
		"code":     200,
		"token":    token,
		"email":    user.Email,
		"number":   user.Number,
		"username": user.UserName,
		"msg":      "登录成功",
	})
}

// Innumber 验证学号是否存在
func Innumber(db *gorm.DB, str string) bool {
	var user model.ComputerUser
	db.Where("number=?", str).Find(&user)
	if len(user.Number) != 0 {
		return true
	}
	return false
}
func InUserName(db *gorm.DB, str string) bool {
	var user model.ComputerUser
	db.Where("user_name=?", str).Find(&user)
	if len(user.UserName) != 0 {
		return true
	}
	return false
}
func InEmail(db *gorm.DB, str string) bool {
	var user model.ComputerUser
	db.Where("email=?", str).Find(&user)
	if len(user.UserName) != 0 {
		return true
	}
	return false
}

// Check MD5加密
//Check 判断是否相等
func Check(content, encrypted string) bool {
	return strings.EqualFold(Encode(content), encrypted)
}

// Encode 加密
func Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// GetUserList 获取用户列表
func GetUserList(g *gin.Context) {
	token := g.Query("token")
	page, _ := strconv.Atoi(g.Query("page"))
	limit, _ := strconv.Atoi(g.Query("limit"))
	data, ok := util.VerifyPermissions(token, g)
	if !ok {
		return
	}
	var UserList []model.ComputerUser
	DB := common.GetDB()
	if data["permissions"] == 1 {
		DB.Select([]string{"id", "name", "profession", "email", "number", "integral"}).Limit(limit).Offset((page - 1) * limit).Find(&UserList)
		g.JSON(200, gin.H{
			"code": 200,
			"data": UserList,
		})
		return
	} else {
		g.JSON(http.StatusForbidden, gin.H{
			"code": 403,
			"data": "未授权访问",
		})
	}

}

// DelUser 删除用户
func DelUser(d *gin.Context) {
	id, _ := strconv.Atoi(d.Query("id"))
	token := d.Query("token")
	data, ok := util.VerifyPermissions(token, d)
	if !ok {
		return
	}
	DB := common.GetDB()
	if data["permissions"] == 1 {
		DB.Delete(model.ComputerUser{}, id)
		d.JSON(200, gin.H{
			"code": 200,
			"msg":  "成功",
		})
		return
	} else {
		d.JSON(http.StatusForbidden, gin.H{
			"code": 403,
			"msg":  "未授权访问",
		})
	}

}

// CheckJwt 验证jwt实效
func CheckJwt(c *gin.Context) {
	token := c.Query("token")
	_, ok := util.VerifyPermissions(token, c)
	if !ok {
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "token正常",
	})
}
