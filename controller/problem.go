package controller

import (
	"CTF/common"
	"CTF/model"
	"CTF/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"os"
)

func ProblemUpload(p *gin.Context) {
	DB := common.GetDB()
	ProblemName := p.PostForm("problem_name")
	ProblemType := p.PostForm("problem_type")
	Answer := p.PostForm("answer")
	Display := p.PostForm("display")
	FlagBool := p.PostForm("flagBool")
	Tips := p.PostForm("tips")
	var Pr model.Problem
	DB.Where("display=?", Display).Find(&Pr)
	if Pr.ID != 0 {
		p.JSON(422, gin.H{
			"msg": "display ID已存在",
		})
		return
	}
	NewProblem := model.Problem{
		ProblemName: ProblemName,
		ProblemType: ProblemType,
		Answer:      Answer,
		Display:     Display,
		Tips:        Tips,
		FlagBool:    FlagBool,
	}
	DB.Create(&NewProblem)
	url := "./data/file/" + ProblemType + "/"

	err := os.MkdirAll(url, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	file, _ := p.FormFile("file")
	fileStr := file.Filename
	for j := len(fileStr); j > 0; j-- {
		if fileStr[j-1] == '.' {
			str := fileStr[j:]
			if str == "zip" || str == "rar" {
				fe := Display + "." + str
				p.SaveUploadedFile(file, url+fe)
				p.JSON(200, gin.H{
					"data": file,
					"msg":  "上传成功",
				})
			} else {
				p.JSON(200, gin.H{
					"data": file,
					"msg":  "文件格式不符合",
				})
				return
			}
		}
	}
}

func MiscProblemDownload(p *gin.Context) {
	id := p.Param("id")
	id = id + ".zip"
	filename := util.LoadFile(id, "./data/file/Misc/")
	fmt.Println(id)
	if filename == "err" {
		p.JSON(500, gin.H{
			"msg": "获取文件失败",
		})
		return
	}
	p.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename)) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	p.Writer.Header().Add("Content-Type", "application/octet-stream")
	p.File("./data/file/Misc/" + filename)
}

// ReverseProblemDownload 逆向题目下载
func ReverseProblemDownload(p *gin.Context) {
	id := p.Param("id")
	id = id + ".zip"
	filename := util.LoadFile(id, "./data/file/Reverse/")
	if filename == "err" {
		p.JSON(500, gin.H{
			"msg": "获取文件失败",
		})
		return
	}
	p.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename)) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	p.Writer.Header().Add("Content-Type", "application/octet-stream")
	p.File("./data/file/Misc/" + filename)
}

// VerifyAnswer 验证答案
func VerifyAnswer(v *gin.Context) {
	type Ans struct {
		Answer   string `json:"answer"`
		UserName string `json:"username"`
		Token    string `json:"token"`
		Number   string `json:"number"`
		Display  string `json:"display"`
	}
	DB := common.GetDB()
	json := Ans{}
	v.BindJSON(&json)
	_, ok := util.VerifyPermissions(json.Token, v)
	if !ok {
		return
	}
	var tmp = model.Problem{}
	var Sub = model.SubmitProblem{}
	DB.Where("display=?", json.Display).Find(&tmp).Update("problem_submit_number", gorm.Expr("problem_submit_number+ ?", 1))
	if json.Answer == tmp.Answer {
		v.JSON(200, gin.H{
			"code": 200,
			"msg":  "答案正确",
		})
		re := common.GetRedis()
		//查找缓存当中是否存在
		w, _ := re.SIsMember(json.Number, json.Display).Result()
		if !w {
			result := DB.Where("number=? AND problem_id=?", json.Number, json.Display).Find(&Sub)
			if result.RowsAffected == 0 {
				//添加到缓存
				re.Do("sadd", json.Number, json.Display)
				//排行榜积分增加
				re.ZIncrBy("Leaderboard", 5, json.UserName)
				//题目通过次数增加
				DB.Where("display=?", json.Display).Find(&tmp).Update("problem_pass_number", gorm.Expr("problem_pass_number+ ?", 1))
				Add := model.SubmitProblem{
					Number:    json.Number,
					ProblemId: json.Display,
				}
				DB.Create(&Add)
			}
		}
	} else {
		v.JSON(400, gin.H{
			"code": 400,
			"msg":  "答案错误",
		})
	}

}
