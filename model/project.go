package model

import "time"

type Project struct {
	UUID      string    `json:"uuid" gorm:"primarykey"`
	Name      string    `json:"name"`
	About     string    `json:"about"`
	Website   string    `json:"website"`
	Repo      string    `json:"repo"`
	Preview   string    `json:"preview"`
	Readme    string    `json:"readme"`
	UID       string    `json:"uid" gorm:"size:191"`
	User      User      `json:"-" gorm:"foreignKey:UID"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (p *Project) TableName() string {
	return "projects"
}
