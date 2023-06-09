package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	// SIGNED_KEY HS256 signed key
	Key = "qwe"
)

// CreateToken @permissions:0代表普通权限，1代表超级用户
func CreateToken(uid string, permissions int) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":         uid,
		"exp":         time.Now().Add(time.Minute * 300).Unix(),
		"permissions": permissions,
	})
	token, err := at.SignedString([]byte(Key))
	if err != nil {
		return "", err
	}

	return token, nil
}

// ParseHStoken 解析签名算法为HS256的token
func ParseHStoken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(Key), nil
	})
	if err != nil {
		return nil, err
	}

	//data：解析包含的过期时间和uid以及用户权限
	data, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("ParseHStoken:claims类型转换失败")
		return nil, err
	}
	return data, nil
}

func VerifyPermissions(token string, v *gin.Context) (jwt.MapClaims, bool) {
	if token == "" {
		v.JSON(403, gin.H{
			"code": 403,
			"data": "token异常",
		})
		return nil, false
	}
	data, err := ParseHStoken(token)
	if err != nil {
		v.JSON(403, gin.H{
			"code": 403,
			"data": "未授权访问",
		})
		return nil, false
	}

	return data, true
}
