package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID uuid.UUID `json:"uuid" gorm:"comment:用户uuid"`
	Username string `json:"username" gorm:"comment:用户登陆名"`
	Password string `json:"password" gorm:"comment:用户登陆密码"`
	NickName string `json:"nick_name" gorm:"comment:用户昵称"`
	HeaderImg string `json:"header_img" gorm:"comment:用户头像"`
}

func (u User) TableName() string {
	return "user"
}