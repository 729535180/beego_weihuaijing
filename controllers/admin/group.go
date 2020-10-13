package admin

import (
	"beego_weihuaijing/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type GroupController struct {
	baseController
}

func (g *GroupController) List() {
	lx := g.Input().Get("type")
	if lx == "json" {
		table := g.o.QueryTable(new(models.Group).TableName()).Filter("status__in", 1, 2)
		name := g.Input().Get("group_name")
		if name != "" {
			table = table.Filter("group_name__icontains", name)
		}

		status, _ := g.GetInt("status")
		if status != 0 {
			table = table.Filter("status__in", status)
		}

		sl, _ := table.Count()
		page := g.PageBase(int(sl))
		var list []models.Group
		table.Limit(page["rows"], page["offset"]).All(&list)

		res := new(models.GroupPage)
		res.Rows = list
		res.Total = int(sl)
		g.Data["json"] = res
		g.ServeJSON()
	} else {
		g.TplName = g.viewsTplName
	}

}

func (g *GroupController) Add() {
	g.TplName = g.controllerName + "/from.html"
}
func (g *GroupController) Edit() {
	lx := g.Input().Get("type")
	if lx == "html" {
		g.TplName = g.controllerName + "/from.html"
	} else {
		id, _ := g.GetInt("id")
		if id == 0 {
			g.Erro("ID值不能为空", "", 0)
		} else {

			cate := models.Group{Id: id}
			g.o.Read(&cate)
			g.Data["json"] = cate
			g.ServeJSON()
		}
	}
	//m.Succ("", "")
}
func (g *GroupController) Save() {
	access := g.GetStrings("ids[]")

	post := models.Group{}
	if len(access) > 0 {
		post.Access = strings.Replace(strings.Trim(fmt.Sprint(access), "[]"), " ", ",", -1)
	}
	post.GroupName = g.Input().Get("group_name")
	post.Status, _ = g.GetInt("status")
	post.UpdateTime = time.Now()
	id, _ := g.GetInt("id")
	if id == 0 {
		post.CreateTime = time.Now()
		if _, err := g.o.Insert(&post); err != nil {
			g.Erro("插入数据错误"+err.Error(), "", 0)
		} else {
			g.Succ("插入数据成功", "")
		}
	} else {
		post.Id = id
		post.CreateTime = time.Now()
		if _, err := g.o.Update(&post); err != nil {
			g.History("更新数据出错"+err.Error(), "")
		} else {
			g.Succ("更新数据成功", "")
		}
	}
}
func (g *GroupController) Del() {
	id := g.Input().Get("id")
	if id == "" {
		g.Erro("ID值不能为空", "", 0)
	} else {
		post := models.Group{}
		post.UpdateTime = time.Now()
		post.Status = 3

		_, err := g.o.QueryTable("tb_group").Filter("id__in", id).Update(orm.Params{
			"status":      3,
			"update_time": time.Now(),
		})
		if err != nil {
			g.History("删除数据出错"+err.Error(), "")
		} else {
			g.Succ("删除数据成功", "")
		}
	}
}
