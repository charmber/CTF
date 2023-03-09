package controller

import (
	"CTF/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// CarouselUpload 轮播图上传
func CarouselUpload(i *gin.Context) {
	err := os.MkdirAll("./data/file/image/carousel", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	file, _ := i.FormFile("file")
	fileStr := file.Filename
	for j := len(fileStr); j > 0; j-- {
		if fileStr[j-1] == '.' {
			str := fileStr[j:]
			if str == "jpg" || str == "png" {
				i.SaveUploadedFile(file, "./data/file/image/carousel/"+file.Filename)
				i.JSON(200, gin.H{
					"data": file,
					"msg":  "上传成功",
				})
			} else {
				i.JSON(200, gin.H{
					"data": file,
					"msg":  "文件格式不符合",
				})
				return
			}
		}
	}

}

// CarouselDownload 获取轮播图
func CarouselDownload(c *gin.Context) {
	id := c.Param("id")
	filename := util.LoadFile(id, "./data/file/image/carousel")
	if filename == "err" {
		c.JSON(500, gin.H{
			"msg": "获取图片失败",
		})
		return
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename)) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("./data/file/image/carousel/" + filename)
}

func LoadImageList(c *gin.Context) {
	FileNameList := util.LoadFileList("./data/file/image/carousel")
	c.JSON(200, gin.H{
		"name": FileNameList,
		"msg":  "获取图片列表成功",
	})
}

// CarouselDel 删除轮播图
func CarouselDel(c *gin.Context) {
	id := c.Param("id")
	if util.DelFile(util.LoadFile(id, "./data/file/image/carousel")) == true {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "删除成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "删除失败，请确认是否存在",
		})
	}
}
