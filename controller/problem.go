package controller

import (
	"CTF/common"
	"CTF/model"
	"CTF/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func ProblemUpload(p*gin.Context)  {
	DB:=common.GetDB()
	ProblemName:=p.PostForm("problem_name")
	ProblemType:=p.PostForm("problem_type")
	Answer:=p.PostForm("answer")
	Display:=p.PostForm("display")
	FlagBool:=p.PostForm("flagBool")
	Tips:=p.PostForm("tips")
	var Pr  model.Problem
	DB.Where("display=?",Display).Find(&Pr)
	if Pr.ID!=0{
		p.JSON(422,gin.H{
			"msg":"display ID已存在",
		})
		return
	}
	NewProblem:=model.Problem{
		ProblemName: ProblemName,
		ProblemType: ProblemType,
		Answer: Answer,
		Display:Display,
		Tips: Tips,
		FlagBool: FlagBool,
	}
	DB.Create(&NewProblem)
	url:="./data/file/"+ProblemType+"/"

	err := os.MkdirAll(url, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	file,_:=p.FormFile("file")
	fileStr:=file.Filename
	for j:=len(fileStr);j>0;j-- {
		if fileStr[j-1] == '.' {
			str := fileStr[j:]
			if str=="zip" || str=="rar"{
				fe:=Display+"."+str
				p.SaveUploadedFile(file,url+fe)
				p.JSON(200,gin.H{
					"data":file,
					"msg":"上传成功",
				})
			}else {
				p.JSON(200,gin.H{
					"data":file,
					"msg":"文件格式不符合",
				})
				return
			}
		}
	}
}

func MiscProblemDownload(p *gin.Context) {
	id:=p.Param("id")
	id=id+".zip"
	filename:=util.LoadFile(id,"./data/file/Misc/")
	fmt.Println(id)
	if filename=="err"{
		p.JSON(500,gin.H{
			"msg":"获取文件失败",
		})
		return
	}
	p.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	p.Writer.Header().Add("Content-Type", "application/octet-stream")
	p.File("./data/file/Misc/"+filename)
	//p.JSON(200,gin.H{
	//	"name":filename,
	//	"msg":"获取文件成功",
	//})
}

// VerifyAnswer 验证答案
func VerifyAnswer(v *gin.Context) {
	ans:=v.Query("ans")
	display:=v.Query("display")
	DB:=common.GetDB()
	var tmp = model.Problem{}
	DB.Where("number=?",display).Find(&tmp)
	if ans==tmp.Answer{
		v.JSON(200,gin.H{
			"msg":"答案正确",
		})
	} else {
		v.JSON(400,gin.H{
			"msg":"答案错误",
		})
	}

}