package initialize

import (
	"github.com/akazwz/go-gin-restful-api/global"
	"github.com/go-redis/redis/v8"
)

func InitRDB() *redis.Client {
	c := global.CFG.RedisDB
	if c.Addr == "" {
		return nil
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password, // no password set
		DB:       0,          // use default DB
	})
	return rdb
}
