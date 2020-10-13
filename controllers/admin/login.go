package admin

import (
	"beego_weihuaijing/models"
	"github.com/astaxie/beego/logs"
)

type LoginController struct {
	baseController
}

func (l *LoginController) Index() {
	if l.Ctx.Request.Method == "POST" {
		username := l.Input().Get("username")
		password := l.Input().Get("password")
		logs.Error("username====", username)
		user := models.AdminUser{Accounts: username}
		l.o.Read(&user, "Accounts")

		if user.Password == "" {
			l.History("账号不存在", "")
		}

		if MyMd5(password) != user.Password {
			l.History("密码错误", "")
		}
		user.LastIp = l.getClientIp()
		user.LoginCount = user.LoginCount + 1
		if _, err := l.o.Update(&user); err != nil {
			l.History("登录异常", "")
		} else {
			l.History("登录成功", "/admin/main.html")
		}
		l.SetSession("admin_user_go", user)
	}
	l.TplName = "admin/login.tpl"
}

func (l *LoginController) Logout() {
	l.DestroySession()
	l.History("退出登录", "/admin/login.html")
}
