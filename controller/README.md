## 支持功能介绍

### article.go

```go
func Recommend(r *gin.Context)
```
主页推荐文件接口，不需要进行身份验证

```go
func CreatArticle(c *gin.Context)
```
创建文章接口，需要进行jwt身份验证

```go
func PageViews(p *gin.Context)
```
文章浏览量增加接口，不需要jwt验证

```go
func ArticleList(g *gin.Context)
```
根据jwt身份查看文章列表

```go
func FindArticle(f *gin.Context)
```
查询搜索文章，根据标题搜索，不需要jwt验证

```go
func FindNotice(f *gin.Context)
```
公告列表前五条，根据创建时间排序，不需要jwt

```go
func DelArticle(d *gin.Context)
```
根据jwt身份进行删除文章


### finproblem.go

```go
func FinProblem(f *gin.Context)
```
获取问题列表，不需要jwt，未登陆状态获取

```go
func UserFinProblem(u *gin.Context)
```
用户登陆状态获取问题列表