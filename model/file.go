package model

import uuid "github.com/satori/go.uuid"

type File struct {
	Model
	UUID     uuid.UUID `json:"uuid"`
	UserUuid string    `json:"user_uuid"`
	URL      string    `json:"url"`
	MD5      string    `json:"md5"`
	Name     string    `json:"name"`
	Size     int64     `json:"size"`
	Type     string    `json:"type"`
}

type FileMD5 struct {
	Model
	MD5      string    `json:"md_5" gorm:"comment:文件MD5"`
	UUID     uuid.UUID `json:"uuid"`
	Location string    `json:"location" gorm:"comment:文件储存位置"`
	Size     int64     `json:"size" gorm:"comment:文件大小"`
	Type     string    `json:"type" gorm:"comment:文件类型"`
	UserUuid string    `json:"user_uuid" gorm:"comment:文件上传者"`
}
