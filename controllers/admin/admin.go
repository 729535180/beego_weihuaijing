package admin

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
	var list []models.Menu
	a.o.QueryTable(new(models.Menu).TableName()).Filter("status__in", 1, 2).Filter("pid", 0).OrderBy("-MenuSort").All(&list)
	a.Data["lists"] = list
	a.Layout = ""
	a.TplName = "admin/index.html"
}

func (a *AdminController) Menu() {
	menuId, _ := a.GetInt("menu")
	//levelId, _ := a.GetInt("levelId")
	var list []models.Menu
	a.o.QueryTable(new(models.Menu).TableName()).Filter("status__in", 1, 2).Filter("pid", menuId).OrderBy("-MenuSort").All(&list)

	a.Data["json"] = list
	a.ServeJSON()
}
