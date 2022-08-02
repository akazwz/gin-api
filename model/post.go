package model

type Post struct {
	UUID    string `json:"uuid" gorm:"primarykey"`
	Title   string `json:"title"`
	Cover   string `json:"cover"`
	Content string `json:"content"`
	Viewed  int    `json:"viewed"`
	UID     string `json:"uid"`
	User    User   `json:"-" gorm:"foreignKey:UID"`
}

func (p *Post) TableName() string {
	return "posts"
}
