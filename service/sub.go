package service

import (
	"github.com/akazwz/go-gin-restful-api/global"
	"github.com/akazwz/go-gin-restful-api/model"
	"github.com/akazwz/go-gin-restful-api/model/request"
	uuid "github.com/satori/go.uuid"
)

func GetSubList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GDB.Model(&model.AllSubWords{})
	var listSub []model.AllSubWords
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&listSub).Error
	return err, listSub, total
}

func GetUserSub(userUuid uuid.UUID) (err error, subI *model.Sub) {
	var sub model.Sub
	err = global.GDB.Where("`user_uuid` = ?", userUuid).First(&sub).Error
	return err, &sub
}

func CreateSub(s *model.Sub) (err error, subI *model.Sub) {
	err = global.GDB.Create(&s).Error
	return err, s
}
