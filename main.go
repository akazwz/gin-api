package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/akazwz/gin-api/global"
	"github.com/akazwz/gin-api/initialize"
	"github.com/joho/godotenv"
)

func main() {
	// 读取环境变量配置
	InitEnvConfig()
	// 初始化路由
	r := initialize.InitRouter()
	// 初始化 gorm
	global.GDB = initialize.InitGorm()
	if global.GDB != nil {
		db, _ := global.GDB.DB()
		err := initialize.RegisterTables(global.GDB)
		if err != nil {
			log.Fatalln("数据库表迁移失败")
		}
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
			}
		}(db)
	} else {
		// gorm 初始化失败
		log.Fatalln("gorm 初始化失败")
	}

	// 端口地址
	port := os.Getenv("API_PORT")
	s := &http.Server{
		Addr:    port,
		Handler: r,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln("Api启动失败")
	}
}

// InitEnvConfig 读取 env 配置文件
func InitEnvConfig() {
	// 非生产环境读取配置文件
	if os.Getenv("GIN_MODE") != "release" {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Fatalln("读取配置文件失败")
		}
	}
}
