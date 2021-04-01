package initialize

import (
	"fmt"
	"github.com/akaedison/go-gin-demo/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//初始化配置
func InitViper() (config *viper.Viper) {
	config = viper.New()
	config.AddConfigPath("./")
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	if err := config.ReadInConfig(); err != nil {
		panic(err)
		return nil
	}

	if err := config.Unmarshal(&global.CFG); err != nil {
		panic(err)
		return nil
	}

	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file updated:", e.Name)
		if err := config.Unmarshal(&global.CFG); err != nil {
			panic(err)
		}
	})

	return
}