package models

import "time"

type Menu struct {
	Id         int       `json:"id"`
	Pid        int       `json:"pid"`
	Text       string    `json:"text"`
	Status     int       `json:"status"`
	IconCls    string    `json:"icon_cls"`
	Domain     string    `json:"domain"`
	Url        string    `json:"url"`
	IsSsl      int       `json:"is_ssl"`
	MenuSort   int       `json:"menu_sort"`
	AdminUid   int       `json:"admin_uid"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (m *Menu) TableName() string {
	return TableName("menu")
}
