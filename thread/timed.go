package thread

import (
	"CTF/common"
	"time"
)

func SyncMysqlRedis() {
	for {
		time.Sleep(time.Millisecond * 100000)
		DB := common.GetDB()
		type Article struct {
			ID int
		}
		var cat []Article
		DB.Select([]string{"id"}).Find(&cat)
		re := common.GetRedis()
		for i := 0; i < len(cat); i++ {
			re.Do("sadd", "article", cat[i].ID)
		}
	}
}
