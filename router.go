package main

import (
	"CTF/controller"
	"CTF/cros"
	"CTF/test"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(cros.Cors()) //跨域
	r.GET("api/test", test.Test)

	//获取轮播图
	r.GET("api/image/carousel/get/:id", controller.CarouselDownload)
	//删除轮播图
	r.GET("api/image/carousel/del/:id", controller.CarouselDel)
	//上传轮播图
	r.POST("api/image/carousel/upload", controller.CarouselUpload)
	//获取图片列表
	r.GET("api/image/carousel/filelist", controller.LoadImageList)

	//文章相关
	r.POST("article/create", controller.CreatArticle)       //添加文章
	r.POST("article/create_notice", controller.CreatNotice) //添加公告
	r.GET("article/get_rec", controller.Recommend)          //推荐文章
	r.GET("article/get_notice", controller.FindNotice)      //获取公告
	r.GET("article/fin", controller.FindArticle)            //查询文章
	r.GET("article/del", controller.DelArticle)             //删除文章
	r.GET("article/list", controller.ArticleList)           //文章列表

	//浏览量
	r.GET("api/article/views/:id", controller.PageViews)

	//用户相关
	r.GET("api/user/getuser", controller.GetUserList) //获取用户列表
	r.GET("api/user/del", controller.DelUser)         //删除一个用户
	r.GET("api/user/token", controller.CheckJwt)      //判断jwt失效

	r.POST("api/user/register", controller.Register)
	r.POST("api/user/login", controller.Login)

	//题目相关

	r.POST("api/problem/upload", controller.ProblemUpload) //上传题目

	//获取问题
	r.GET("api/problem/:id", controller.FinProblem) //根据分类获取题目列表,未登录获取列表

	r.POST("api/problem/:id", controller.UserFinProblem) //登录获取列表

	r.GET("api/problem/misc/:id", controller.MiscProblemDownload)       //杂项问题下载接口
	r.GET("api/problem/reverse/:id", controller.ReverseProblemDownload) //逆向问题下载接口
	//验证答案
	r.POST("api/problem/answer", controller.VerifyAnswer)

	//排行榜系统
	r.GET("api/leaderboard/find", controller.InquireLeaderboard)

	return r
}
