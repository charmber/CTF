package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	id:=c.Param("id")
	filename:=id+".jpg"
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("./file/qqqq.jpg")
}
func Test01(c *gin.Context) {
	filename:="a.txt"
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("./file/a.txt")
	fmt.Println("a.txt")
}
func Test02(c *gin.Context) {
	filename:="归档.zip"
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("./file/归档.zip")
	fmt.Println("zip")
}

func Tes(t *gin.Context) {
	t.JSON(200,gin.H{
		"code":200,
		"data":"你好，第一个接口",
		"msg":"发布成功",
	})
}
