package service

import (
	"github.com/akazwz/go-gin-demo/global"
	"github.com/akazwz/go-gin-demo/model"
	uuid "github.com/satori/go.uuid"
)

// CreateUserFile create user file
func CreateUserFile(f *model.File) (err error) {
	f.UUID = uuid.NewV4()
	err = global.GDB.Create(&f).Error
	return err
}

// CreateMD5File create md5 file
func CreateMD5File(f5 *model.FileMD5) (err error) {
	f5.UUID = uuid.NewV4()
	err = global.GDB.Create(&f5).Error
	return err
}
