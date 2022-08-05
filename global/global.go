package global

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	GDB    *gorm.DB
	GREDIS *redis.Client
	R2C    *s3.Client
)
