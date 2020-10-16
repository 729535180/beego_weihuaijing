package admin

import (
	"beego_weihuaijing/models"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
	baseController
}

func (u *UserController) GroupData() error {
	var list []models.Group
	u.o.QueryTable(new(models.Group).TableName()).Filter("status__in", 1).All(&list)
	u.Data["group_list"] = list

	return nil
}

func (u *UserController) Add() {
	if err := u.GroupData(); err != nil {
		u.Erro("读取数据出错"+err.Error(), "", 0, u.resData)
	}
	u.Data["password"] = true
	u.TplName = u.controllerName + "/from.html"
}
func (u *UserController) Edit() {
	lx := u.Input().Get("type")
	if lx == "html" {
		if err := u.GroupData(); err != nil {
			u.Erro("读取数据出错"+err.Error(), "", 0, u.resData)
		}
		u.Data["password"] = false
		u.TplName = u.controllerName + "/from.html"
	} else {
		id, _ := u.GetInt("id")
		if id == 0 {
			u.Erro("ID值不能为空", "", 0, u.resData)
		} else {

			cate := models.AdminUser{Id: id}
			u.o.Read(&cate)
			cate.Password = ""
			u.Data["json"] = cate
			u.ServeJSON()
		}
	}
	//m.Succ("", "")
}
func (u *UserController) Save() {
	id, _ := u.GetInt("id")
	post := models.AdminUser{}
	post.Accounts = u.Input().Get("accounts")
	post.Username = u.Input().Get("username")
	post.Level, _ = u.GetInt("level")
	post.Status, _ = u.GetInt("status")
	post.UpdateTime = time.Now()
	groupId := u.GetStrings("group_id")
	if len(groupId) > 0 {
		post.GroupId = strings.Replace(strings.Trim(fmt.Sprint(groupId), "[]"), " ", ",", -1)
	}
	post.Password = u.Input().Get("password")
	if post.Password != "" {
		post.Password = MyMd5(post.Password)
	}

	if id == 0 {
		post.CreateTime = time.Now()
		if _, err := u.o.Insert(&post); err != nil {
			u.Erro("插入数据错误"+err.Error(), "", 0, u.resData)
		} else {
			u.Succ("插入数据成功", "", u.resData)
		}
	} else {
		post.Id = id
		post.UpdateTime = time.Now()
		if post.Password == "" {
			if _, err := u.o.Update(&post, "Accounts", "Username", "GroupId", "Level", "Status", "UpdateTime"); err != nil {
				u.Erro("更新数据出错"+err.Error(), "", 0, u.resData)
			} else {
				u.Succ("更新数据成功", "", u.resData)
			}
		} else {
			if _, err := u.o.Update(&post, "Accounts", "Password", "Username", "GroupId", "Level", "Status", "UpdateTime"); err != nil {
				u.Erro("更新数据出错"+err.Error(), "", 0, u.resData)
			} else {
				u.Succ("更新数据成功", "", u.resData)
			}
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
		u.Erro("ID值不能为空", "", 0, u.resData)
	} else {
		post := models.Menu{}
		post.UpdateTime = time.Now()
		post.Status = 3
		post.Id, _ = strconv.Atoi(id)
		if _, err := u.o.Update(&post, "Status", "UpdateTime"); err != nil {
			u.Erro("删除数据出错"+err.Error(), "", 0, u.resData)
		} else {
			u.Succ("删除数据成功", "", u.resData)
		}

	}
}
