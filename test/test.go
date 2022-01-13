package test

import (
	"CTF/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Test(t *gin.Context) {
	rdb := common.GetRedis()
	w, _ := rdb.SIsMember("20200909", "8").Result()
	fmt.Println(w)
}
