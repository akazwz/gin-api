package model

import (
	"time"
)

type User struct {
	UUID     string `json:"uuid" gorm:"primarykey"`
	Username string `json:"username"`
	/* 返回 json 时忽略 password 字段 */
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (u User) TableName() string {
	return "users"
}
