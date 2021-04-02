package service

import (
	"github.com/akaedison/go-gin-demo/global"
	"github.com/akaedison/go-gin-demo/model"
)

func Login(u *model.User) (err error, userInter *model.User) {
	var user model.User
	err = global.GDB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}
