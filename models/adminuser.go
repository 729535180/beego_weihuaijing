package models

import (
	"time"
)

type AdminUser struct {
	Id         int       `json:"id"`
	Accounts   string    `json:"accounts"`
	Password   string    `json:"password"`
	Username   string    `json:"username"`
	GroupId    string    `json:"group_id"`
	Level      int       `json:"level"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	LastIp     string    `json:"last_ip"`
	LoginCount int       `json:"login_count"`
}

type AdminUserPage struct {
	Rows  []AdminUser `json:"rows"`
	Total int         `json:"total"`
}

func (m *AdminUser) TableName() string {
	return TableName("admin_user")
}
