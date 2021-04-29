package model

import uuid "github.com/satori/go.uuid"

type Book struct {
	Model
	UUID         uuid.UUID `json:"uuid" gorm:"comment:图书uuid"`
	BookName     string    `json:"book_name" gorm:"comment:图书名"`
	Author       string    `json:"author" gorm:"comment:作者"`
	Price        float64   `json:"price" gorm:"comment:价格"`
	Cover        string    `json:"cover" gorm:"comment:封面图片"`
	Introduction string    `json:"introduction" gorm:"comment: 简介"`
}

func (b Book) TableName() string {
	return "book"
}
