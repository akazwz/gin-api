package model

import (
	"time"
)

type Model struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"index"`
}
