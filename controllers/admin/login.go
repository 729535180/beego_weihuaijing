package admin

import (
	"beego_weihuaijing/models"
)

type LoginController struct {
	baseController
}

func (l *LoginController) Index() {
	if l.Ctx.Request.Method == "POST" {
		username := l.Input().Get("username")
		password := l.Input().Get("password")

		user := models.AdminUser{Accounts: username}
		l.o.Read(&user, "Accounts")

		if user.Password == "" {
			l.Erro("账号不存在", "操作失败", 0, l.resData)
		}

		if MyMd5(password) != user.Password {
			l.Erro("密码错误", "操作失败", 0, l.resData)
		}
		user.LastIp = l.getClientIp()

		user.LoginCount = user.LoginCount + 1
		if _, err := l.o.Update(&user, "LastIp", "LoginCount"); err != nil {
			l.Erro("登录异常", "操作失败", 0, l.resData)
		} else {
			l.Succ("登录成功", "操作成功", l.resData)
		}
		l.SetSession("admin_user_go", user)

	}
	l.TplName = "admin/login.tpl"
}

func (l *LoginController) Logout() {
	l.DestroySession()
	l.History("退出登录", "/admin/login.html")
}
