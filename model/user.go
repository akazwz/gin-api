package model

import (
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Model
	UUID        uuid.UUID `json:"uuid" gorm:"unique comment:用户uuid"`
	Username    string    `json:"username" gorm:"unique comment:用户登陆名"`
	Password    string    `json:"password" gorm:"comment:用户登陆密码"`
	Phone       string    `json:"phone" gorm:"unique comment:手机号"`
	OpenId      string    `json:"open_id" gorm:"unique comment:微信OpenId"`
	AuthorityId string    `json:"authority_id" gorm:"default:000;comment:用户角色ID"`
	NickName    string    `json:"nick_name" gorm:"comment:用户昵称"`
	AvatarUrl   string    `json:"avatar_url" gorm:"default:https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTJ6m9WsdeiclaLlZ9YIibCYZ7HlS8NroTvnDlLvj91BIju5UB36I51Px0uLMpVahrwRq8Mk1t2mSCRQ/132;comment:用户头像"`
	Gender      int       `json:"gender" gorm:"comment:用户性别"`
	Bio         string    `json:"bio" gorm:"comment:用户签名"`
}

func (u User) TableName() string {
	return "user"
}
