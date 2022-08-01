package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	UUID     uuid.UUID `json:"uuid" gorm:"unique"`
	Username string    `json:"username"`
	/* 返回 json 时忽略 password 字段 */
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (u User) TableName() string {
	return "users"
}
