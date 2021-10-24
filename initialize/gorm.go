package initialize

import (
	"fmt"
	"os"

	"github.com/akazwz/go-gin-restful-api/global"
	"github.com/akazwz/go-gin-restful-api/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB 初始化数据库连接
func InitDB() *gorm.DB {
	m := global.CFG.Database

	if m.Name == "" {
		return nil
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		m.User,
		m.Password,
		m.Host,
		m.Name,
	)

	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return nil
	} else {
		//sqlDB, _ := db.DB()
		//sqlDB.SetMaxIdleConns()
		//sqlDB.SetMaxIdleConns()
		return db
	}
}

func CreateTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
		model.Book{},
		model.File{},
		model.FileMD5{},
		model.Sub{},
		model.AllSubWords{},
	)
	if err != nil {
		os.Exit(0)
	}
}
