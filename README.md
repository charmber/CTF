## 基于GO-Gin+Grom的后端服务框架

适用于初步入门，快速搭建后端的新手

只需要下载完成后，初始化go.mod自动下载相应依赖包即可

在控制台执行

`go run main.go routes.go`命令即可完成部署





### 目录介绍

**common**

存放数据库相关配置文件以及相关启动接口函数

**controller**

后端接口函数

**cros**

跨域解决函数

**model**

数据库创建字段存放结构体

**response**

数据发送接收请求函数封装

**util**

自定义函数存放目录

**`go.mod`和`go.sum`**

包管理

`**main.go**`

存放主函数

`**routes.go**`

路由文件

### 特别功能介绍

在接收发送请求方面，采用目录基本主流的大部分方式，例如POST，GET，DELETE，PUT；

在接收数据方面，有采用**表单**，**JSON**，**PARAMS**等主流的传输方式

登录密码采用MD5加密

查询数据有进行分页处理

对后端进行了跨域问题的解决

### 该框架配合系统

**后端管理后台Vue+Gin+Grom**

[Vue后台开源地址](https://github.com/charmber/Vue_element-Backstage.git)

**小程序管理程序+Gin+Grom**

[小程序开源地址](https://github.com/charmber/WeChat-Mini-Program.git)


