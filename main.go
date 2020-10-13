package main

import (
	"beego_weihuaijing/models"
	_ "beego_weihuaijing/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	models.Init()
	orm.Debug = true
	beego.BConfig.WebConfig.Session.SessionOn = true
}
func main() {
	beego.Run()
}
