package model

import uuid "github.com/satori/go.uuid"

type HotSearch struct {
	Model
	UUID      uuid.UUID `json:"uuid"`
	Rank      int64
	Content   string
	Link      string
	Hot       int64
	TopicLead string
}

func (hot HotSearch) TableName() string {
	return "hot_search"
}
