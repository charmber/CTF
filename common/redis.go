package common

import (
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func InitRedis() *redis.Client {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 8,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		panic("failed to connect redis,err:" + err.Error())
	}
	return rdb
}

func GetRedis() *redis.Client {
	RDB := rdb
	return RDB
}
