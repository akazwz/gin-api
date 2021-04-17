package service

import (
	"github.com/akaedison/go-gin-demo/global"
	"github.com/akaedison/go-gin-demo/model"
	uuid "github.com/satori/go.uuid"
)

func AddBook(b *model.Book) (err error, bookInter *model.Book) {
	b.UUID = uuid.NewV4()
	err = global.GDB.Create(&b).Error
	return err, b
}

func DeleteBook(id float64) (err error) {
	var book model.Book
	err = global.GDB.Where("id = ?", id).Delete(&book).Error
	return err
}
