## 基于GO-Gin+Gorm的CTF后端服



运行`go build main.go routes.go`命令进行构建


### 目录介绍

**common**

存放数据库相关配置文件以及相关启动接口函数

**controller**

业务逻辑函数

**cros**

跨域解决函数

**model**

数据库创建字段存放结构体


**util**

自定义函数存放目录

`go.mod`和`go.sum`

包管理

`main.go`

存放主函数

`routes.go`

路由文件

### 介绍


整个项目采用微服务架构，主模块与沙箱环境模块，两个模块之间使用gRPC进行通信

登录JWT身份令牌

单独开启一个线程使用redis作为存入数据库的缓冲

使用docker作为比赛环境以及靶机的沙箱




[CTF前端页面开源地址](https://github.com/charmber/Vue_element-Backstage.git)


