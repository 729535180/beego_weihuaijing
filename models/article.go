package models

import "time"

type Article struct {
	Id          int       `json:"id"`
	ClassifyId  int       `json:"classify_id"`
	Title       string    `json:"title"`
	Keywords    string    `json:"keywords"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Ip          string    `json:"ip"`
	Read        int       `json:"read"`
	Send_time   time.Time `json:"send_time"`
	Author      string    `json:"author"`
	Source      string    `json:"source"`
	Url         string    `json:"url"`
	Img         string    `json:"img"`
	TagIds      string    `json:"tag_ids"`
	Sort        int       `json:"sort"`
	Status      int       `json:"status"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
}
type ArticlePage struct {
	Rows  []Article `json:"rows"`
	Total int       `json:"total"`
}

func (a *Article) TableName() string {
	return TableName("article")
}
