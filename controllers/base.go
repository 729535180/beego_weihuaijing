package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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
	b.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	b.actionName = strings.ToLower(actionName)
	b.viewsTplName = b.controllerName + "/" + b.actionName
	b.o = orm.NewOrm()
	logs.Error("controller====", b.controllerName)
	logs.Error("action====", b.actionName)

}
