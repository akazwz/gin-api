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
	//counter.Consume()
	// 读取环境变量配置
	InitEnvConfig()

	// 初始化路由
	r := initialize.InitRouter()
	// 初始化 R2
	global.R2C = initialize.InitR2Client()
	// 初始化 gorm
	global.GDB = initialize.InitGorm()
	db, _ := global.GDB.DB()
	initialize.RegisterTables(global.GDB)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	// 初始化 redis
	global.GREDIS = initialize.InitRedis()

	// 端口地址
	port := os.Getenv("API_PORT")
	s := &http.Server{
		Addr:    port,
		Handler: r,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
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
