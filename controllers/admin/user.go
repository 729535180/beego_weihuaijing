package admin

import (
	"beego_weihuaijing/models"
	"time"
)

type UserController struct {
	baseController
}

func (u *UserController) Add() {
	u.TplName = u.controllerName + "/from.html"
}
func (u *UserController) Save() {
	post := models.AdminUser{}
	post.Accounts = u.Input().Get("accounts")
	post.Password = u.Input().Get("password")
	post.Username = u.Input().Get("username")
	post.Level, _ = u.GetInt("level")
	post.Status, _ = u.GetInt("status")
	post.UpdateTime = time.Now()
	id, _ := u.GetInt("id")
	if id == 0 {
		post.CreateTime = time.Now()
		if _, err := u.o.Insert(&post); err != nil {
			u.Erro("插入数据错误"+err.Error(), "", 0)
		} else {
			u.Succ("插入数据成功", "")
		}
	} else {
		post.Id = id
		post.CreateTime = time.Now()
		if _, err := u.o.Update(&post); err != nil {
			u.History("更新数据出错"+err.Error(), "")
		} else {
			u.Succ("更新数据成功", "")
		}
	}
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
