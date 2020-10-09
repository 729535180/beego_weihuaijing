package admin

import (
	"beego_weihuaijing/models"
)

type UserController struct {
	baseController
}

func (u *UserController) List() {
	lx := u.Input().Get("type")
	if lx == "json" {
		sl, _ := u.o.QueryTable(new(models.AdminUser).TableName()).Filter("status__in", 1, 2).Count()
		page := u.PageBase(int(sl))
		var list []models.AdminUser
		u.o.QueryTable(new(models.AdminUser).TableName()).Filter("status__in", 1, 2).Limit(page["rows"], page["offset"]).All(&list)

		res := new(models.AdminUserPage)
		res.Rows = list
		res.Total = int(sl)
		u.Data["json"] = res
		u.ServeJSON()
		//logs.Error("view====", sl)
	} else {
		u.TplName = u.viewsTplName
	}
	//logs.Error("view====", u.viewsTplName)

}
