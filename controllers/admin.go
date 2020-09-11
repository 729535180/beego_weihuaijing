package controllers

import (
	"beego_weihuaijing/models"
	"github.com/astaxie/beego/logs"
)

type AdminController struct {
	baseController
}
type Menu struct {
	Id      int    `json:"id"`
	Pid     int    `json:"pid"`
	Text    string `json:"text"`
	State   string `json:"state"`
	IconCls string `json:"iconCls"`
	Url     string `json:"url"`
}

func (a *AdminController) Index() {
	logs.Error("views====", a.viewsTplName)
	a.TplName = "admin/index.tpl"
}

func (a *AdminController) Menu() {
	menuId, _ := a.GetInt("menu")
	levelId, _ := a.GetInt("levelId")
	logs.Error("id======", menuId)
	logs.Error("levelId======", levelId)
	var list []models.Menu
	a.o.QueryTable(new(models.Menu).TableName()).All(&list)

	a.Data["json"] = list
	a.ServeJSON()
}
