package service

import (
	"github.com/akazwz/go-gin-restful-api/global"
	"github.com/akazwz/go-gin-restful-api/model"
	"github.com/akazwz/go-gin-restful-api/model/request"
	uuid "github.com/satori/go.uuid"
)

func CreateBook(b *model.Book) (err error, bookInter *model.Book) {
	b.UUID = uuid.NewV4()
	err = global.GDB.Create(&b).Error
	return err, b
}

func DeleteBook(id float64) (err error) {
	var book model.Book
	err = global.GDB.Where("id = ?", id).Delete(&book).Error
	return err
}

func UpdateBook(b *model.Book) (err error, book *model.Book) {
	err = global.GDB.Updates(&b).Error
	return err, b
}

func GetBookList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GDB.Model(&model.Book{})
	var bookList []model.Book
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&bookList).Error
	return err, bookList, total
}
