package routers

import (
	"beego_weihuaijing/controllers"
	"beego_weihuaijing/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin/login/index", &admin.LoginController{}, "*:Index")
	beego.Router("/admin/index", &admin.AdminController{}, "*:Index")
	beego.Router("/admin/menu", &admin.AdminController{}, "*:Menu")
	beego.Router("/admin/upload", &admin.AdminController{}, "*:Upload")
	ns := beego.NewNamespace("/admin",
		// 后台页面
		beego.NSRouter("/", &admin.AdminController{}, "*:Index"),
		beego.NSRouter("/index", &admin.AdminController{}, "*:Index"),
		beego.NSRouter("/menu", &admin.AdminController{}, "*:Menu"),

		// 管理员列表
		beego.NSNamespace("/user",
			beego.NSRouter("/", &admin.UserController{}, "*:List"),
			beego.NSRouter("/list", &admin.UserController{}, "*:List"),
			beego.NSRouter("/add", &admin.UserController{}, "*:Add"),
			beego.NSRouter("/save", &admin.UserController{}, "*:Save"),
			beego.NSRouter("/edit", &admin.UserController{}, "*:Edit"),
			beego.NSRouter("/del", &admin.UserController{}, "*:Del"),
		),
		// 菜单列表
		beego.NSNamespace("/menu",
			beego.NSRouter("/", &admin.MenuController{}, "*:List"),
			beego.NSRouter("/list", &admin.MenuController{}, "*:List"),
			beego.NSRouter("/menu_json", &admin.MenuController{}, "*:DataJson"),
			beego.NSRouter("/add", &admin.MenuController{}, "*:Add"),
			beego.NSRouter("/edit", &admin.MenuController{}, "*:Edit"),
			beego.NSRouter("/save", &admin.MenuController{}, "*:Save"),
			beego.NSRouter("/del", &admin.MenuController{}, "*:Del"),
		),
		//角色组
		beego.NSNamespace("/group",
			beego.NSRouter("/", &admin.GroupController{}, "*:List"),
			beego.NSRouter("/list", &admin.GroupController{}, "*:List"),
			beego.NSRouter("/add", &admin.GroupController{}, "*:Add"),
			beego.NSRouter("/save", &admin.GroupController{}, "*:Save"),
			beego.NSRouter("/edit", &admin.GroupController{}, "*:Edit"),
			beego.NSRouter("/del", &admin.GroupController{}, "*:Del"),
		), //beego.NSRouter("/", &controllers.AdminGroupController{},"*:Index"),
		//beego.NSRouter("/create", &controllers.AdminGroupController{},"*:Create"),

		//管理员日志
		beego.NSNamespace("/log"), //beego.NSRouter("/", &controllers.AdminLogController{},"*:Index"),

		//权限管理
		beego.NSNamespace("/rule"), //beego.NSRouter("/", &controllers.AdminRuleController{},"*:Index"),

		//CMS 分类
		beego.NSNamespace("/classify",
			beego.NSRouter("/", &admin.ClassifyController{}, "*:List"),
			beego.NSRouter("/list", &admin.ClassifyController{}, "*:List"),
			beego.NSRouter("/add", &admin.ClassifyController{}, "*:Add"),
			beego.NSRouter("/save", &admin.ClassifyController{}, "*:Save"),
			beego.NSRouter("/edit", &admin.ClassifyController{}, "*:Edit"),
			beego.NSRouter("/del", &admin.ClassifyController{}, "*:Del"),
		),
		//CMS 文章
		beego.NSNamespace("/article",
			beego.NSRouter("/", &admin.ArticleController{}, "*:List"),
			beego.NSRouter("/list", &admin.ArticleController{}, "*:List"),
			beego.NSRouter("/add", &admin.ArticleController{}, "*:Add"),
			beego.NSRouter("/save", &admin.ArticleController{}, "*:Save"),
			beego.NSRouter("/edit", &admin.ArticleController{}, "*:Edit"),
			beego.NSRouter("/del", &admin.ArticleController{}, "*:Del"),
		),
	)
	beego.AddNamespace(ns)
	//beego.AutoRouter(&admin.AdminController{})
}
