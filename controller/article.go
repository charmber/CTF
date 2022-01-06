package controller

import (
	"CTF/common"
	"CTF/model"
	"CTF/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Recommend 推荐文章
func Recommend(r*gin.Context)  {
	DB:=common.GetDB()
	var cat []model.Article
	DB.Limit(5).Order("recommend desc").Find(&cat)
	r.JSON(http.StatusOK,gin.H{
		"code":0,
		"msg":"",
		"len":"5",
		"count":1000,
		"data":cat,
	})
}

// CreatArticle 创建文章
func CreatArticle(c *gin.Context) {
	DB:=common.GetDB()
	type Article struct {
		User	  string `json:"user"`  //账号
		Name      string `json:"name"`   //用户名
		Title     string `json:"title"`  //标题
		NewsType  string `json:"news_type"`   //文章类型
		Author    string `json:"author"`   //作者
		Content   string `json:"content"`          //文章内容
		Time      string `json:"time"`   //时间
	}
	token:=c.Query("token")
	if !util.VerifyPermissions(token,c){
		return
	}
	json:=Article{}
	c.BindJSON(&json)
	NewArt:=model.Article{
		User: json.User,
		Name: json.Name,
		Title: json.Title,
		NewsType: json.NewsType,
		Author: json.Author,
		Content: json.Content,
		Time: json.Time,
	}
	DB.Create(&NewArt)
	//返回结果
	c.JSON(200,gin.H{
		"code":200,
		"data":gin.H{
			"data":NewArt},
		"msg":"发布成功",
	})
}

// PageViews 文章浏览量
func PageViews(p*gin.Context)  {
	id:=p.Param("id")
	DB:=common.GetDB()
	var art model.Article
	DB.Where("id=?",id).Find(&art).Update("Recommend",art.Recommend+1)
	p.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"",
		"count":1000,
		"data":art,
	})
}
func PageViewsArticle(id int)  {

}

// FindArticle 查询文章
func FindArticle(f*gin.Context){
	DB:=common.GetDB()
	title:=f.Query("title")
	var cat model.Article
	DB.Where("title=?",title).Find(&cat)
	f.JSON(200,gin.H{
		"data":cat,
		"msg":"查询成功",
	})


}
