package admin

import "github.com/astaxie/beego/logs"

type UserController struct {
	baseController
}

func (u *UserController) List() {
	logs.Error("view====", u.viewsTplName)
}
