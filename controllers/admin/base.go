package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
