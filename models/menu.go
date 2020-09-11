package models

type Menu struct {
	Id             int
	Pid            int
	Text           string
	Status         int
	IconCls        string
	Domain         string
	MenuModule     string
	MenuController string
	MenuAction     string
	Param          string
	IsSsl          int
	MenuSort       int
	AdminUid       int
	CreateTime     string
	UpdateTime     string
}

func (m *Menu) TableName() string {
	return TableName("menu")
}
