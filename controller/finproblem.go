package controller

import (
	"CTF/common"
	"CTF/model"
	"CTF/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FinProblem(f *gin.Context) {
	DB := common.GetDB()
	var tmp []model.Problem
	var problem []model.Problem
	id := f.Param("id")

	DB.Where("problem_type=?", id).Scopes(util.Paginate(f)).Find(&tmp)
	DB.Find(&problem)
	f.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取成功",
		"len":  len(problem),
		"data": tmp,
	})
}

func UserFinProblem(u *gin.Context) {
	type Problem struct {
		Number string `json:"number"`
		Token  string `json:"token"`
	}
	id := u.Param("id")
	json := Problem{}
	u.BindJSON(&json)
	DB := common.GetDB()
	re := common.GetRedis()
	var tmp []model.Problem
	var problem []model.Problem
	DB.Where("problem_type=?", id).Scopes(util.Paginate(u)).Find(&tmp)
	DB.Find(&problem)
	r, err := re.SMembers(json.Number).Result()
	if err != nil {
		fmt.Println("get abc faild :", err)
		return
	}
	u.JSON(http.StatusOK, gin.H{
		"code":    0,
		"msg":     "获取成功",
		"len":     len(problem),
		"data":    tmp,
		"through": r,
	})
}
