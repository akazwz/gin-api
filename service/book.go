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
