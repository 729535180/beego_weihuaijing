package admin

import (
	"beego_weihuaijing/models"
	"strconv"
	"time"
)

type ClassifyController struct {
	baseController
}

func (c *ClassifyController) List() {
	pId, _ := c.GetInt("id")
	lx := c.Input().Get("type")
	if lx == "" {
		c.TplName = c.viewsTplName
	} else {
		table := c.o.QueryTable(new(models.Classify).TableName())
		var list []models.Classify
		table.Filter("status__in", 1, 2).Filter("pid", pId).OrderBy("-Sort").All(&list)
		var newListArray []*models.Classify
		if pId == 0 && lx == "of" {
			newList := new(models.Classify)
			newList.Id = 0
			newList.Pid = 0
			newList.Name = "一级分类"
			newList.Text = "一级分类"
			newList.Des = ""
			newList.Img = ""
			newList.Sort = 0
			newList.Status = 1
			newList.State = "open"
			newListArray = append(newListArray, newList)
		}
		for _, v := range list {
			nrData := new(models.Classify)
			nrData.Id = v.Id
			nrData.Pid = v.Pid
			nrData.Name = v.Name
			nrData.Text = v.Name
			nrData.Des = v.Des
			nrData.Img = v.Img
			nrData.Sort = v.Sort
			nrData.Status = v.Status
			nrData.State = v.State
			num, _ := table.Filter("status__in", 1, 2).Filter("pid", v.Id).Count()
			if num > 0 {
				nrData.State = "closed"
			} else {
				nrData.State = ""
			}
			newListArray = append(newListArray, nrData)
		}

		if len(newListArray) < 1 {

			xv := [][]int64{}
			c.Data["json"] = xv
		} else {
			c.Data["json"] = newListArray
		}

		c.ServeJSON()
	}
}

func (c *ClassifyController) Add() {
	c.TplName = c.controllerName + "/from.html"
}

func (c *ClassifyController) Edit() {
	lx := c.Input().Get("type")
	if lx == "html" {
		c.TplName = c.controllerName + "/from.html"
	} else {
		id, _ := c.GetInt("id")
		if id == 0 {
			c.Erro("ID值不能为空", "", 0, c.resData)
		} else {

			cate := models.Classify{Id: id}
			c.o.Read(&cate)
			c.Data["json"] = cate
			c.ServeJSON()
		}
	}

}

func (c *ClassifyController) Save() {
	post := models.Classify{}

	post.Name = c.Input().Get("name")
	post.Des = c.Input().Get("des")
	post.Img = c.Input().Get("img")
	post.Pid, _ = c.GetInt("pid")
	post.Sort, _ = c.GetInt("sort")
	post.Status, _ = c.GetInt("status")
	post.UpdateTime = time.Now()
	id, _ := c.GetInt("id")
	if id == 0 {
		post.CreateTime = time.Now()
		if _, err := c.o.Insert(&post); err != nil {
			c.Erro("插入数据错误"+err.Error(), "", 0, c.resData)
		} else {
			c.Succ("插入数据成功", "", c.resData)
		}
	} else {
		post.Id = id
		if _, err := c.o.Update(&post, "Name", "Des", "Img", "Pid", "Sort", "Status", "UpdateTime"); err != nil {
			c.History("更新数据出错"+err.Error(), "")
		} else {
			c.Succ("更新数据成功", "", c.resData)
		}
	}
}

func (c *ClassifyController) Del() {
	id := c.Input().Get("id")
	if id == "" {
		c.Erro("ID值不能为空", "", 0, c.resData)
	} else {
		post := models.Classify{}
		post.UpdateTime = time.Now()
		post.Status = 3
		post.Id, _ = strconv.Atoi(id)

		if _, err := c.o.Update(&post, "Status", "UpdateTime"); err != nil {
			c.History("删除数据出错"+err.Error(), "")
		} else {
			c.Succ("删除数据成功", "", c.resData)
		}

	}
}
