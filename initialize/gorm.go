package initialize

import (
	"github.com/akazwz/gin-api/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitGorm() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}

func RegisterTables(db *gorm.DB) error {
	err := db.AutoMigrate(
		model.User{},
	)
	return err
}
