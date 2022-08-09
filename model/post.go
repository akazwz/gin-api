package model

import (
	"time"
)

type Post struct {
	UUID      string    `json:"uuid" gorm:"primarykey"`
	Title     string    `json:"title"`
	Cover     string    `json:"cover"`
	Content   string    `json:"content"`
	UID       string    `json:"uid" gorm:"size:191"`
	User      User      `json:"-" gorm:"foreignKey:UID"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (p *Post) TableName() string {
	return "posts"
}
