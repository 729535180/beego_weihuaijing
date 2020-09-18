package admin

import "beego_weihuaijing/models"

type MenuController struct {
	baseController
}
type MenuJson struct {
	Id      int    `json:"id"`
	Pid     int    `json:"pid"`
	Text    string `json:"text"`
	State   string `json:"state"`
	IconCls string `json:"iconCls"`
	Url     string `json:"url"`
	Uid     int    `json:"uid"`
	Sort    int    `json:"sort"`
	Status  int    `json:"status"`
}

func (m *MenuController) List() {
	//logs.Error("view====", m.viewsTplName)
	m.TplName = m.viewsTplName
	//m.Layout = m.Layout
}
func (m *MenuController) DataJson() {
	menuId, _ := m.GetInt("id")
	var list []models.Menu
	m.o.QueryTable(new(models.Menu).TableName()).Filter("status__in", 1, 2).Filter("pid", menuId).OrderBy("-MenuSort").All(&list)
	var newListArray []*MenuJson
	for _, v := range list {
		newList := new(MenuJson)
		newList.Id = v.Id
		newList.Uid = v.Id
		newList.Pid = v.Pid
		newList.Text = v.Text
		newList.IconCls = v.IconCls
		newList.Url = v.Url
		newList.Sort = v.MenuSort
		newList.Status = v.Status
		if v.Url == "" {
			newList.State = "closed"
		} else {
			newList.State = "open"
		}
		newListArray = append(newListArray, newList)
	}
	if len(newListArray) < 1 {
		newList := new(MenuJson)
		newList.Id = 0
		newList.Pid = 0
		newList.Text = "无数据"
		newList.IconCls = ""
		newList.Url = "无数据"
		newList.State = "open"
		newListArray = append(newListArray, newList)
	}
	m.Data["json"] = newListArray
	m.ServeJSON()

}
func (m *MenuController) Add() {
	m.TplName = m.controllerName + "/from.html"
}
func (m *MenuController) Edit() {
	m.Succ("", "")
}

func (m *MenuController) Save() {
	m.Succ("", "")
}
