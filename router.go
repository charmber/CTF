package main

import (
	"CTF/controller"
	"CTF/cros"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(cros.Cors())//跨域

	//获取轮播图
	r.GET("api/image/carousel/get/:id",controller.CarouselDownload)
	//删除轮播图
	r.GET("api/image/carousel/del/:id",controller.CarouselDel)
	//上传轮播图
	r.POST("api/image/carousel/upload",controller.CarouselUpload)

	r.GET("file/down/2",controller.Test01)
	r.GET("file/down/3",controller.Test02)

	r.GET("test",controller.Tes)
	r.POST("article/create",controller.CreatArticle)		//添加文章
	r.GET("article/get_rec",controller.Recommend)		//推荐文章
	r.GET("article/fin",controller.FindArticle)			//查询文章
	//浏览量
	r.GET("api/article/views/:id",controller.PageViews)

	r.POST("api/user/register",controller.Register)
	r.POST("api/user/login",controller.Login)


	//题目相关
	r.POST("api/problem/upload",controller.ProblemUpload)
	r.GET("api/problem",controller.FinProblem)
	r.GET("api/problem/misc/:id",controller.MiscProblemDownload) //杂项问题下载接口
	return r
}
