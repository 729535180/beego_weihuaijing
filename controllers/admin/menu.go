package admin

import (
	"beego_weihuaijing/models"
	"time"
)

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

		xv := [][]int64{}
		m.Data["json"] = xv
	} else {
		m.Data["json"] = newListArray
	}

	m.ServeJSON()

}
func (m *MenuController) Add() {
	m.TplName = m.controllerName + "/from.html"
}
func (m *MenuController) Edit() {
	lx := m.Input().Get("type")
	if lx == "html" {
		m.TplName = m.controllerName + "/from.html"
	} else {
		id, _ := m.GetInt("id")
		if id == 0 {
			m.Erro("ID值不能为空", "", 0)
		} else {

			cate := models.Menu{Id: id}
			m.o.Read(&cate)
			m.Data["json"] = cate
			m.ServeJSON()
		}
	}
	//m.Succ("", "")
}

func (m *MenuController) Save() {
	post := models.Menu{}
	post.Text = m.Input().Get("text")
	post.Url = m.Input().Get("url")
	post.IconCls = m.Input().Get("icon_cls")
	post.Status, _ = m.GetInt("status")
	post.MenuSort, _ = m.GetInt("menu_sort")

	post.UpdateTime = time.Now()

	id, _ := m.GetInt("id")
	if id == 0 {
		post.CreateTime = time.Now()
		if _, err := m.o.Insert(&post); err != nil {
			m.Erro("插入数据错误"+err.Error(), "", 0)
		} else {
			m.Succ("插入数据成功", "")
			//m.History("插入数据成功", "/admin/index.html")
		}
	} else {
		post.Id = id
		post.CreateTime = time.Now()
		if _, err := m.o.Update(&post); err != nil {
			m.History("更新数据出错"+err.Error(), "")
		} else {
			m.Succ("更新数据成功", "")
			//m.History("插入数据成功", "/admin/index.html")
		}
	}

}

func (m *MenuController) Del() {
	id, _ := m.GetInt("id")
	if id == 0 {
		m.Erro("ID值不能为空", "", 0)
	} else {
		post := models.Menu{}
		post.UpdateTime = time.Now()
		post.Status = 3
		post.Id = id
		if _, err := m.o.Update(&post, "Status", "UpdateTime"); err != nil {
			m.History("删除数据出错"+err.Error(), "")
		} else {
			m.Succ("删除数据成功", "")
			//m.History("插入数据成功", "/admin/index.html")
		}
	}
}
