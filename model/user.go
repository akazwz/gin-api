package model

import (
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Model
	UUID        uuid.UUID `json:"uuid" gorm:"comment:用户uuid"`
	Username    string    `json:"username" gorm:"comment:用户登陆名"`
	Password    string    `json:"password" gorm:"comment:用户登陆密码"`
	NickName    string    `json:"nick_name" gorm:"comment:用户昵称"`
	HeaderImg   string    `json:"header_img" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	AuthorityId string    `json:"authority_id" gorm:"default:000;comment:用户角色ID"`
}

func (u User) TableName() string {
	return "user"
}
