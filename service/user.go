package service

import (
	"errors"
	"github.com/akaedison/go-gin-demo/global"
	"github.com/akaedison/go-gin-demo/model"
	"github.com/akaedison/go-gin-demo/pkg/util"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func Register(u model.User) (err error, userInter model.User) {
	var user model.User
	if !errors.Is(global.GDB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("username already exist"), userInter
	}
	// md5 register
	u.Password = util.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.GDB.Create(&u).Error
	return err, u
}

func Login(u *model.User) (err error, userInter *model.User) {
	var user model.User
	u.Password = util.MD5V([]byte(u.Password))
	err = global.GDB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}

func FindUserByID(id int) (err error, user *model.User) {
	var u model.User
	err = global.GDB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

func FindUserByUUID(uuid string) (err error, user *model.User) {
	var u model.User
	if err = global.GDB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("user does not exist"), &u
	}
	return nil, &u
}
