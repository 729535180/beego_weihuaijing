package models

import "time"

type Classify struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Des        string    `json:"des"`
	Pid        int       `json:"pid"`
	Img        string    `json:"img"`
	Sort       int       `json:"sort"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	State      string    `json:"state" orm:"-"`
	Text       string    `json:"text" orm:"-"`
}

func (c *Classify) TableName() string {
	return TableName("classify")
}
