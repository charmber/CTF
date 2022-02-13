package test

import (
	"CTF/model/global"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Test(f *gin.Context) {
	id := f.Query("id")
	global.Operate <- 1
	global.DockerID <- id
	fmt.Println("test success!")
}
