package controller

import (
	"CTF/common"
	"CTF/model"
	"CTF/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"strconv"
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
	_, ok := util.VerifyPermissions(token, c)
	if !ok {
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

// ArticleList 查看自己的文章列表or管理员查看所有人文章
func ArticleList(g *gin.Context) {
	DB := common.GetDB()
	var cat []model.Article
	token := g.Query("token")
	data, ok := util.VerifyPermissions(token, g)
	if !ok {
		return
	}
	if cast.ToInt(data["permissions"]) == 0 {
		DB.Where("number=?", cast.ToString(data["uid"])).Scopes(util.Paginate(g)).Find(&cat)
		g.JSON(200, gin.H{
			"data": cat,
			"msg":  "查询成功",
		})
	} else if data["permissions"] == 1 {
		DB.Scopes(util.Paginate(g)).Find(&cat)
		g.JSON(200, gin.H{
			"data": cat,
			"msg":  "查询成功",
		})
	} else {
		g.JSON(403, gin.H{
			"data": cat,
			"msg":  "未授权访问",
		})
	}
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

// FindNotice 查找公告
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

// DelArticle 删除文章
func DelArticle(d *gin.Context) {
	DB := common.GetDB()
	token := d.Query("token")
	data, ok := util.VerifyPermissions(token, d)
	if !ok {
		return
	}
	var tmp model.Article
	id, _ := strconv.Atoi(d.Query("id"))

	//判断用户权限
	if data["permissions"] == 0 {
		DB.Where("id=?", id).Find(&tmp)
		if tmp.Number == data["uid"] {
			DB.Delete(model.Article{}, id)
			d.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "删除成功",
			})
			return
		} else {
			d.JSON(http.StatusForbidden, gin.H{
				"code": 403,
				"msg":  "未授权访问",
			})
			return
		}
	} else if data["permissions"] == 1 {
		DB.Delete(model.Article{}, id)
		d.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "删除成功",
		})
		return
	}

}
