package admin

import (
	"beego_weihuaijing/models"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ArticleController struct {
	baseController
}

func (a *ArticleController) GroupData() error {
	var list []models.Group
	a.o.QueryTable(new(models.Group).TableName()).Filter("status__in", 1).All(&list)
	a.Data["group_list"] = list

	return nil
}

func (a *ArticleController) Add() {
	if err := a.GroupData(); err != nil {
		a.Erro("读取数据出错"+err.Error(), "", 0, a.resData)
	}
	a.Data["password"] = true
	a.TplName = a.controllerName + "/from.html"
}
func (a *ArticleController) Edit() {
	lx := a.Input().Get("type")
	if lx == "html" {
		if err := a.GroupData(); err != nil {
			a.Erro("读取数据出错"+err.Error(), "", 0, a.resData)
		}
		a.Data["password"] = false
		a.TplName = a.controllerName + "/from.html"
	} else {
		id, _ := a.GetInt("id")
		if id == 0 {
			a.Erro("ID值不能为空", "", 0, a.resData)
		} else {

			cate := models.AdminUser{Id: id}
			a.o.Read(&cate)
			cate.Password = ""
			a.Data["json"] = cate
			a.ServeJSON()
		}
	}
	//m.Succ("", "")
}
func (a *ArticleController) Save() {
	id, _ := a.GetInt("id")
	post := models.AdminUser{}
	post.Accounts = a.Input().Get("accounts")
	post.Username = a.Input().Get("username")
	post.Level, _ = a.GetInt("level")
	post.Status, _ = a.GetInt("status")
	post.UpdateTime = time.Now()
	groupId := a.GetStrings("group_id")
	if len(groupId) > 0 {
		post.GroupId = strings.Replace(strings.Trim(fmt.Sprint(groupId), "[]"), " ", ",", -1)
	}
	post.Password = a.Input().Get("password")
	if post.Password != "" {
		post.Password = MyMd5(post.Password)
	}

	if id == 0 {
		post.CreateTime = time.Now()
		if _, err := a.o.Insert(&post); err != nil {
			a.Erro("插入数据错误"+err.Error(), "", 0, a.resData)
		} else {
			a.Succ("插入数据成功", "", a.resData)
		}
	} else {
		post.Id = id
		post.UpdateTime = time.Now()
		if post.Password == "" {
			if _, err := a.o.Update(&post, "Accounts", "Username", "GroupId", "Level", "Status", "UpdateTime"); err != nil {
				a.Erro("更新数据出错"+err.Error(), "", 0, a.resData)
			} else {
				a.Succ("更新数据成功", "", a.resData)
			}
		} else {
			if _, err := a.o.Update(&post, "Accounts", "Password", "Username", "GroupId", "Level", "Status", "UpdateTime"); err != nil {
				a.Erro("更新数据出错"+err.Error(), "", 0, a.resData)
			} else {
				a.Succ("更新数据成功", "", a.resData)
			}
		}
	}
}
func (a *ArticleController) List() {
	lx := a.Input().Get("type")
	if lx == "json" {
		table := a.o.QueryTable(new(models.Article).TableName()).Filter("status__in", 1, 2)
		title := a.Input().Get("title")
		if title != "" {
			table = table.Filter("title__icontains", title)
		}

		level, _ := a.GetInt("level")
		if level != 0 {
			table = table.Filter("level__in", level)
		}
		status, _ := a.GetInt("status")
		if status != 0 {
			table = table.Filter("status__in", status)
		}

		sl, _ := table.Count()
		page := a.PageBase(int(sl))
		var list []models.Article
		table.Limit(page["rows"], page["offset"]).All(&list)

		res := new(models.ArticlePage)
		res.Rows = list
		res.Total = int(sl)
		a.Data["json"] = res
		a.ServeJSON()
	} else {
		a.TplName = a.viewsTplName
	}

}
func (a *ArticleController) Del() {
	id := a.Input().Get("id")
	if id == "" {
		a.Erro("ID值不能为空", "", 0, a.resData)
	} else {
		post := models.Menu{}
		post.UpdateTime = time.Now()
		post.Status = 3
		post.Id, _ = strconv.Atoi(id)
		if _, err := a.o.Update(&post, "Status", "UpdateTime"); err != nil {
			a.Erro("删除数据出错"+err.Error(), "", 0, a.resData)
		} else {
			a.Succ("删除数据成功", "", a.resData)
		}

	}
}
