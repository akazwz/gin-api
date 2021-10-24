package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/datatypes"
)

type Sub struct {
	Model
	UserUUID uuid.UUID      `json:"user_uuid"`
	SubWords datatypes.JSON `json:"sub_words"`
}

type AllSubWords struct {
	Model
	SubWord string `json:"sub_word" gorm:"unique"`
}

func (s Sub) TableName() string {
	return "sub"
}

func (a AllSubWords) TableName() string {
	return "all_sub_words"
}
