package models

import "time"

type AdminUser struct {
	Id         int       `json:"id"`
	Accounts   string    `json:"accounts"`
	Password   string    `json:"password"`
	Username   string    `json:"username"`
	Level      int       `json:"level"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
type AdminUserPage struct {
	Rows  []AdminUser `json:"rows"`
	Total int         `json:"total"`
}

func (m *AdminUser) TableName() string {
	return TableName("admin_user")
}
