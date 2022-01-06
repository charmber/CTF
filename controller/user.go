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
	"strings"
)

func Register(c *gin.Context) {

	DB:=common.GetDB()
	//获取参数
	/*number:=lg.PostForm("number")
	password:=lg.PostForm("password")*/
	type Username struct {
		UserName   string `json:"user_name"` //用户名账号
		Email      string `json:"email"` //邮箱
		NickName   string `json:"nick_name"`        //昵称
		Name       string `json:"name"`        //姓名
		Number     string `json:"number"`         //学号
		Password   string `json:"password"`       //密码
		Academy    string `json:"academy"`            //学院
		Profession string `json:"profession"`            //专业
		CreateTime string `json:"create_time"`
	}
	json :=Username{}
	c.BindJSON(&json)

	//数据验证
	if len(json.UserName)==0{
		c.JSON(422,gin.H{
			"code":422,
			"msg":"用户名账号必须填写(英文/数字/符号)",
		})
	}
	if len(json.Email)==0{
		c.JSON(422,gin.H{
			"code":422,
			"msg":"邮箱账号必须填写",
		})
	}
	if len(json.Name)==0{
		c.JSON(422,gin.H{
			"code":422,
			"msg":"用户名账号必须填写(英文/数字/符号)",
		})
	}
	if len(json.Password)<=6{
		c.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"密码不能少于6位"})
		return
	}
	//如果人名为空，就随机一个10位字符串
	if len(json.NickName)==0{
		json.Name=util.Random(10)
	}

	//验证是否用户存在
	if InUserName(DB,json.UserName){
		c.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"用户名已经存在"})
		return
	}
	if InEmail(DB,json.Email){
		c.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"邮箱用户已经存在"})
		return
	}
	if len(json.Number)!=0{
		if Innumber(DB,json.Number){
			c.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"学号已经存在"})
			return
		}
	}

	//创建用户
	//加密密码
	md:=Encode(json.Password)
	newUser:=model.ComputerUser{
		Name: json.Name,
		Number: json.Number,
		Password: md,
		UserName: json.UserName,
		Email: json.Email,
		NickName: json.NickName,
		Academy: json.Academy,
		Profession: json.Profession,
		CreateTime: json.CreateTime,
	}
	DB.Create(&newUser)


	//返回结果
	c.JSON(200,gin.H{
		"code":200,
		"msg":"注册成功",
	})
}


func Login(lg *gin.Context) {

	DB:=common.GetDB()
	//获取参数
	/*number:=lg.PostForm("number")
	password:=lg.PostForm("password")*/
	//tp:1账号，2用户名，3邮箱
	type Username struct {
		Tp string `json:"tp"`
		Number string `json:"number"`
		Password string `json:"password"`
	}
	var user model.ComputerUser
	json :=Username{}
	lg.BindJSON(&json)
	if json.Tp=="1"{
		//数据验证
		if len(json.Password)<=6{
			lg.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"密码不能少于6位"})
			return
		}
		//判断是否存在
		DB.Where("number=?",json.Number).Find(&user)
		if user.ID==0{
			lg.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"用户不存在"})
			return
		}
	}else if json.Tp=="2" {
		//数据验证
		if len(json.Password)<=6{
			lg.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"密码不能少于6位"})
			return
		}
		//判断是否存在
		DB.Where("user_name=?",json.Number).Find(&user)
		if user.ID==0{
			lg.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"用户不存在"})
			return
		}
	}else if json.Tp=="3" {
		//数据验证
		if len(json.Password)<=6{
			lg.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"密码不能少于6位"})
			return
		}
		//判断是否存在
		DB.Where("email=?",json.Number).Find(&user)
		if user.ID==0{
			lg.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"用户不存在"})
			return
		}
	}else {
		lg.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"不存在登录方式"})
	}


	if Check(json.Password,user.Password)!=true{
		lg.JSON(http.StatusBadRequest,gin.H{"code":400,"msg":"密码错误"})
		return
	}
	//if json.Password!=user.Password{
	//	lg.JSON(http.StatusBadRequest,gin.H{"code":400,"msg":"密码错误"})
	//		return
	//}


	//发放token
	token,_:=util.CreateToken(json.Number)

	//返回结果
	lg.JSON(200,gin.H{
		"code":200,
		"data":gin.H{"token":token,
			"nickname":user.NickName},
		"msg":"登录成功",
	})
}

// Innumber 验证学号是否存在
func Innumber(db *gorm.DB, str string) bool{
	var user model.ComputerUser
	db.Where("number=?",str).Find(&user)
	if len(user.Number)!=0{
		return true
	}
	return false
}
func InUserName(db *gorm.DB,str string) bool {
	var user model.ComputerUser
	db.Where("user_name=?",str).Find(&user)
	if len(user.UserName)!=0{
		return true
	}
	return false
}
func InEmail(db *gorm.DB,str string) bool {
	var user model.ComputerUser
	db.Where("email=?",str).Find(&user)
	if len(user.UserName)!=0{
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
