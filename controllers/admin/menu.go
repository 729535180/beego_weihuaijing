package admin

import (
	"beego_weihuaijing/models"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type MenuController struct {
	baseController
}
type MenuJson struct {
	Id      int    `json:"id"`
	Ids     int    `json:"ids[]"`
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
	lx := m.Input().Get("type")
	var list []models.Menu
	m.o.QueryTable(new(models.Menu).TableName()).Filter("status__in", 1, 2).Filter("pid", menuId).OrderBy("-MenuSort").All(&list)
	var newListArray []*MenuJson
	if menuId == 0 && lx == "of" {
		newList := new(MenuJson)
		newList.Id = 0
		newList.Ids = 0
		newList.Uid = 0
		newList.Pid = 0
		newList.Text = "主菜单"
		newList.IconCls = ""
		newList.Url = ""
		newList.Sort = 0
		newList.Status = 1
		newList.State = "open"
		newListArray = append(newListArray, newList)
	}
	for _, v := range list {
		newList := new(MenuJson)
		newList.Id = v.Id
		newList.Ids = v.Id
		newList.Uid = v.Id
		newList.Pid = v.Pid
		newList.Text = v.Text
		newList.IconCls = v.IconCls
		newList.Url = v.Url
		newList.Sort = v.MenuSort
		newList.Status = v.Status
		if v.Url == "" {
			if lx == "of" {
				if newList.Pid > 0 {
					newList.State = "closed"
				} else {
					newList.State = "closed"
				}
			} else {
				newList.State = "closed"
			}
		} else {
			newList.State = "closed"
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
	post.Pid, _ = m.GetInt("pid")

	post.UpdateTime = time.Now()
	if post.Pid > 0 {
		cate := models.Menu{Id: post.Pid}
		m.o.Read(&cate, "id")
		if cate.Id != 0 {
			if cate.Path == "0" {
				post.Path = strconv.Itoa(cate.Id)
			} else {
				post.Path = cate.Path + "," + strconv.Itoa(cate.Id)
			}
		}
	} else {
		post.Path = "0"
	}

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
	id := m.Input().Get("id")
	if id == "" {
		m.Erro("ID值不能为空", "", 0)
	} else {
		post := models.Menu{}
		post.UpdateTime = time.Now()
		post.Status = 3

		_, err := m.o.QueryTable("tb_menu").Filter("id__in", id).Update(orm.Params{
			"status":      3,
			"update_time": time.Now(),
		})
		if err != nil {
			m.History("删除数据出错"+err.Error(), "")
		} else {
			m.Succ("删除数据成功", "")
			//m.History("插入数据成功", "/admin/index.html")
		}
	}
}
