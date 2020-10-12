package models

import (
	"time"
)

type Group struct {
	Id         int       `json:"id"`
	GroupName  string    `json:"group_name"`
	Access     string    `json:"access"`
	AdminUid   int       `json:"admin_uid"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
type GroupPage struct {
	Rows  []Group `json:"rows"`
	Total int     `json:"total"`
}

func (g *Group) TableName() string {
	return TableName("group")
}
