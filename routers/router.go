package routers

import (
	"beego_weihuaijing/controllers"
	"beego_weihuaijing/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin/index", &admin.AdminController{}, "*:Index")
	beego.Router("/admin/menu", &admin.AdminController{}, "*:Menu")
	ns := beego.NewNamespace("/admin",
		// 后台页面
		beego.NSRouter("/", &admin.AdminController{}, "*:Index"),
		beego.NSRouter("/index", &admin.AdminController{}, "*:Index"),
		beego.NSRouter("/menu", &admin.AdminController{}, "*:Menu"),

		// 管理员列表
		beego.NSNamespace("/user",
			beego.NSRouter("/", &admin.UserController{}, "*:List"),
			beego.NSRouter("/list", &admin.UserController{}, "*:List"),
		),
		// 菜单列表
		beego.NSNamespace("/menu",
			beego.NSRouter("/", &admin.MenuController{}, "*:List"),
			beego.NSRouter("/list", &admin.MenuController{}, "*:List"),
			beego.NSRouter("/menu_json", &admin.MenuController{}, "*:DataJson"),
			beego.NSRouter("/add", &admin.MenuController{}, "*:Add"),
			beego.NSRouter("/save", &admin.MenuController{}, "*:Save"),
		),
		//角色组
		beego.NSNamespace("/group"), //beego.NSRouter("/", &controllers.AdminGroupController{},"*:Index"),
		//beego.NSRouter("/create", &controllers.AdminGroupController{},"*:Create"),

		//管理员日志
		beego.NSNamespace("/log"), //beego.NSRouter("/", &controllers.AdminLogController{},"*:Index"),

		//权限管理
		beego.NSNamespace("/rule"), //beego.NSRouter("/", &controllers.AdminRuleController{},"*:Index"),

	)
	beego.AddNamespace(ns)
	//beego.AutoRouter(&admin.AdminController{})
}
