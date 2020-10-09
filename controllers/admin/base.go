package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"strings"
)

type baseController struct {
	beego.Controller
	controllerName string
	actionName     string
	viewsTplName   string
	o              orm.Ormer
}

func (b *baseController) Prepare() {
	controllerName, actionName := b.GetControllerAndAction()
	b.controllerName = "admin/" + strings.ToLower(controllerName[0:len(controllerName)-10])
	b.actionName = strings.ToLower(actionName)
	b.viewsTplName = b.controllerName + "/" + b.actionName + ".html"
	b.Layout = "admin/layout.html"
	b.o = orm.NewOrm()
	b.TplName = b.viewsTplName
	//logs.Error("controller====", b.controllerName)
	//logs.Error("action====", b.actionName)

}

type comReturn struct {
	StatusCode int    `json:"statusCode"`
	Title      string `json:"title"`
	Message    string `json:"message"`
}

//topjui  操作成功返回
func (b *baseController) Succ(msg string, title string) {
	re := new(comReturn)
	if msg == "" {
		re.Message = "恭喜你，操作成功！"
	} else {
		re.Message = msg
	}
	if title == "" {
		re.Title = "操作提示"
	} else {
		re.Title = title
	}
	re.StatusCode = 200
	b.Data["json"] = re
	b.ServeJSON()
}

//topjui  操作失败返回
func (b *baseController) Erro(msg string, title string, code int) {
	re := new(comReturn)
	if msg == "" {
		re.Message = "操作失败！"
	} else {
		re.Message = msg
	}
	if title == "" {
		re.Title = "操作提示"
	} else {
		re.Title = title
	}
	if code == 0 {
		re.StatusCode = 404
	} else {
		re.StatusCode = code
	}

	b.Data["json"] = re
	b.ServeJSON()
}

//日志信息输出到页面上
func (b *baseController) History(msg string, url string) {
	if url == "" {
		b.Ctx.WriteString("<script>alert('" + msg + "');window.history.go(-1);</script>")
		b.StopRun()
	} else {
		b.Redirect(url, 302)
	}
}

/*
page 当前页数
rows 每页数量
nums 数据总量
*/
func (b *baseController) PageBase(nums int) map[string]interface{} {

	var offset, totalpage, lastpage, firstpage, next, prev, page, rows int
	if page, _ = b.GetInt("page"); page < 1 {
		page = 1
	}
	if rows, _ = b.GetInt("rows"); rows < 1 {
		rows = 20
	}
	totalpage = int(math.Ceil(float64(nums / rows)))
	lastpage = totalpage
	firstpage = 1

	offset = (page - 1) * rows
	if (page + 1) > totalpage {
		next = totalpage
	} else {
		next = page + 1
	}
	if (page - 1) <= 0 {
		prev = 1
	} else {
		prev = page - 1
	}
	pageMap := make(map[string]interface{})
	pageMap["totalpage"] = totalpage //总页数
	pageMap["lastpage"] = lastpage   //最后一页
	pageMap["firstpage"] = firstpage //最开始一页
	pageMap["page"] = page           //当前页
	pageMap["rows"] = rows           //每页数量
	pageMap["offset"] = offset       //数据库开始位置
	pageMap["next"] = next           //下一页
	pageMap["prev"] = prev           //上一页

	return pageMap
}
