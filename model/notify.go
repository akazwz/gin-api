package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Notify struct {
	Model
	UserUUID         uuid.UUID `json:"user_uuid" gorm:"user_uuid comment:用户uuid"`
	NotifyGap        int       `json:"notify_gap" gorm:"notify_gap comment:通知间隔,单位分钟"`
	NotifyCount      int       `json:"notify_count" gorm:"notify_count comment:通知次数"`
	AllNotifyCount   int       `json:"all_notify_count" gorm:"all_notify_count comment:总通知次数"`
	LastNotify       time.Time `json:"last_notify" gorm:"last_notify comment:上次发送通知日期"`
	NotifyLimitCount int       `json:"notify_limit_count" gorm:"notify_limit_count"`
}

func (n Notify) TableName() string {
	return "notify"
}
