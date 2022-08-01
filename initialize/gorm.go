package initialize

import (
	"log"

	"github.com/akazwz/gin-api/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitGorm() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("初始化 gorm 失败")
	}
	return db
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
	)
	if err != nil {
		log.Fatalln("数据库表迁移失败")
	}
}
