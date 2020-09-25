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
		var (
			page   int
			rows   int = 20
			offset int
		)
		if page, _ = u.GetInt("page"); page < 1 {
			page = 1
		}
		if rows, _ = u.GetInt("rows"); rows < 1 {
			rows = 20
		}
		offset = (page - 1) * rows
		var list []models.AdminUser
		u.o.QueryTable(new(models.AdminUser).TableName()).Filter("status__in", 1, 2).Limit(rows, offset).All(&list)
		sl, _ := u.o.QueryTable(new(models.AdminUser).TableName()).Filter("status__in", 1, 2).Count()
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
