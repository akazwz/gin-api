package initialize

import (
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

func InitRedis() *redis.Client {
	options, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Println(err)
		log.Fatalln("解析redis url 失败")
	}
	return redis.NewClient(options)
}
