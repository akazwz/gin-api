package global

import (
	"github.com/akazwz/go-gin-restful-api/config"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GDB  *gorm.DB
	GRDB *redis.Client
	VP   *viper.Viper
	CFG  config.Conf
	_    *zap.Logger
)
