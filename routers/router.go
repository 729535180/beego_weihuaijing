package routers

import (
	"beego_weihuaijing/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.AdminController{}, "*:Index")

	beego.AutoRouter(&controllers.AdminController{})
}
