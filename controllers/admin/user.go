package admin

import (
	"beego_weihuaijing/models"
	"github.com/astaxie/beego/orm"
	"time"
)

type UserController struct {
	baseController
}

func (u *UserController) Add() {
	u.TplName = u.controllerName + "/from.html"
}
func (u *UserController) Edit() {
	lx := u.Input().Get("type")
	if lx == "html" {
		u.TplName = u.controllerName + "/from.html"
	} else {
		id, _ := u.GetInt("id")
		if id == 0 {
			u.Erro("ID值不能为空", "", 0)
		} else {

			cate := models.AdminUser{Id: id}
			u.o.Read(&cate)
			u.Data["json"] = cate
			u.ServeJSON()
		}
	}
	//m.Succ("", "")
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
		table := u.o.QueryTable(new(models.AdminUser).TableName()).Filter("status__in", 1, 2)
		name := u.Input().Get("username")
		if name != "" {
			table = table.Filter("username__icontains", name)
		}
		acco := u.Input().Get("accounts")
		if acco != "" {
			table = table.Filter("accounts__icontains", acco)
		}
		level, _ := u.GetInt("level")
		if level != 0 {
			table = table.Filter("level__in", level)
		}
		status, _ := u.GetInt("status")
		if status != 0 {
			table = table.Filter("status__in", status)
		}

		sl, _ := table.Count()
		page := u.PageBase(int(sl))
		var list []models.AdminUser
		table.Limit(page["rows"], page["offset"]).All(&list)

		res := new(models.AdminUserPage)
		res.Rows = list
		res.Total = int(sl)
		u.Data["json"] = res
		u.ServeJSON()
	} else {
		u.TplName = u.viewsTplName
	}

}
func (u *UserController) Del() {
	id := u.Input().Get("id")
	if id == "" {
		u.Erro("ID值不能为空", "", 0)
	} else {
		post := models.Menu{}
		post.UpdateTime = time.Now()
		post.Status = 3

		_, err := u.o.QueryTable("tb_admin_user").Filter("id__in", id).Update(orm.Params{
			"status":      3,
			"update_time": time.Now(),
		})
		if err != nil {
			u.History("删除数据出错"+err.Error(), "")
		} else {
			u.Succ("删除数据成功", "")
		}
	}
}
