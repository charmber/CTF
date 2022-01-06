package util

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"math/rand"
	"strconv"
	"time"
)

// Random 随机字符串
func Random(n int) string {
	var letter =[]byte("qwertyuiopasdfghjklzxcvbnm");
	result:=make([]byte,n)
	rand.Seed(time.Now().Unix())
	for i:=range result {
		result[i]=letter[rand.Intn(len(letter))]
	}
	return string(result)
}

// Paginate get分页查询函数
func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.Query("limit"))

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)		//offset开始页数  limit返回条数
	}
}