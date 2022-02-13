package controller

import (
	"CTF/common"
	"github.com/gin-gonic/gin"
)

func InquireLeaderboard(i *gin.Context) {
	re := common.GetRedis()
	ans, _ := re.ZRevRange("Leaderboard", 0, 10).Result()
	i.JSON(200, gin.H{
		"data": ans,
	})
}
