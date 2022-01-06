package controller

import (
	"CTF/common"
	"CTF/model"
	"CTF/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FinProblem(f *gin.Context)  {
	DB:=common.GetDB()
	var tmp []model.Problem
	var problem []model.Problem
	f.Get("/")
	DB.Scopes(util.Paginate(f)).Find(&tmp)
	DB.Find(&problem)
	f.JSON(http.StatusOK,gin.H{
		"code":0,
		"msg":"获取成功",
		"len":len(problem),
		"data":tmp,
	})
}
