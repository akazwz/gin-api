package utils

import (
	"context"
	"encoding/json"
	"time"

	"github.com/akazwz/gin-api/global"
)

func RedisCacheSet(key string, value interface{}, ttl time.Duration) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = global.GREDIS.Set(context.Background(), key, string(bytes), ttl).Err()
	if err != nil {
		return err
	}
	return nil
}

func RedisCacheGet(key string, result interface{}) error {
	res, err := global.GREDIS.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(res), result)
	if err != nil {
		return err
	}
	return nil
}

func RedisCacheDel(key string) error {
	return global.GREDIS.Del(context.Background(), key).Err()
}
