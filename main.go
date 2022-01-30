package main

import (
	"CTF/common"
	"CTF/thread"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	db := common.InitDB()
	rdb := common.InitRedis()
	go thread.SyncMysqlRedis()

	defer db.Close()

	defer rdb.Close()

	r := gin.Default()
	r = CollectRouter(r)
	panic(r.Run(":9000"))
}

//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go router.go
