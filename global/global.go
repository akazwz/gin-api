package global

import (
	"github.com/akazwz/go-gin-demo/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GDB *gorm.DB
	VP  *viper.Viper
	CFG config.Conf
	LOG *zap.Logger
)
