package controller

import (
	"CTF/common"
	"CTF/model"
	"CTF/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Recommend 推荐文章
func Recommend(r *gin.Context) {
	DB := common.GetDB()
	var cat []model.Article
	DB.Limit(5).Order("recommend desc").Find(&cat)
	r.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "",
		"len":   "5",
		"count": 1000,
		"data":  cat,
	})
}

// CreatArticle 创建文章
func CreatArticle(c *gin.Context) {
	DB := common.GetDB()
	type Article struct {
		Number   string `json:"number"` //账号
		UserName string `json:"username"`
		Name     string `json:"name"`      //名字
		Title    string `json:"title"`     //标题
		NewsType string `json:"news_type"` //文章类型
		Author   string `json:"author"`    //作者
		Content  string `json:"content"`   //文章内容
		Time     string `json:"time"`      //时间
	}
	token := c.Query("token")
	if !util.VerifyPermissions(token, c) {
		return
	}
	json := Article{}
	c.BindJSON(&json)
	NewArt := model.Article{
		Number:   json.Number,
		Name:     json.Name,
		Title:    json.Title,
		NewsType: json.NewsType,
		Author:   json.Author,
		Content:  json.Content,
		Time:     json.Time,
	}
	DB.Create(&NewArt)
	re := common.GetRedis()
	re.Do("sadd", "article", NewArt.ID)
	re.ZIncrBy("Leaderboard", 1, json.UserName)
	//返回结果
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"data": NewArt},
		"msg": "发布成功",
	})
}

func CreatNotice(c *gin.Context) {
	DB := common.GetDB()
	type Notice struct {
		Name    string `json:"name"`    //用户名
		Title   string `json:"title"`   //标题
		Author  string `json:"author"`  //作者
		Content string `json:"content"` //文章内容
		Time    string `json:"time"`    //时间
	}
	token := c.Query("token")
	fmt.Println("qwe")
	if !util.VerifyPermissions(token, c) {
		return
	}
	fmt.Println("测试")
	json := Notice{}
	c.BindJSON(&json)
	NewArt := model.Notice{
		Name:    json.Name,
		Title:   json.Title,
		Author:  json.Author,
		Content: json.Content,
		Time:    json.Time,
	}
	DB.Create(&NewArt)
	//返回结果
	c.JSON(200, gin.H{
		"code": 200,
		"data": NewArt,
		"msg":  "发布成功",
	})
}

// PageViews 文章浏览量
func PageViews(p *gin.Context) {
	id := p.Param("id")
	re := common.GetRedis()
	w, _ := re.SIsMember("article", id).Result() //先查询缓存是否存在
	if !w {
		p.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "未查询到该文章",
		})
		return
	}
	DB := common.GetDB()
	var art model.Article
	DB.Where("id=?", id).Find(&art).Update("Recommend", art.Recommend+1)
	p.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "",
		"count": 1000,
		"data":  art,
	})
}
func PageViewsArticle(id int) {

}

// FindArticle 查询文章
func FindArticle(f *gin.Context) {
	DB := common.GetDB()
	title := f.Query("title")
	var cat model.Article
	DB.Where("title=?", title).Find(&cat)
	f.JSON(200, gin.H{
		"data": cat,
		"msg":  "查询成功",
	})
}

func FindNotice(f *gin.Context) {
	DB := common.GetDB()
	type Notice struct {
		Time  string
		Title string
	}
	var cat []Notice
	DB.Limit(5).Order("created_at desc").Select([]string{"time", "title"}).Find(&cat)
	f.JSON(200, gin.H{
		"data": cat,
	})
}
