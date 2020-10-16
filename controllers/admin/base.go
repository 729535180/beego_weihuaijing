package admin

import (
	"beego_weihuaijing/common"
	"beego_weihuaijing/models"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"strconv"
	"strings"
)

type baseController struct {
	beego.Controller
	controllerName string
	actionName     string
	viewsTplName   string
	rouleName      string
	o              orm.Ormer
	loginUser      models.AdminUser
	showMenuId     string
	resData        map[string]string
}

func (b *baseController) Prepare() {
	controllerName, actionName := b.GetControllerAndAction()
	b.controllerName = "admin/" + strings.ToLower(controllerName[0:len(controllerName)-10])
	b.actionName = strings.ToLower(actionName)
	b.viewsTplName = b.controllerName + "/" + b.actionName + ".html"
	b.rouleName = strings.ToLower(controllerName[0:len(controllerName)-10]) + "/" + b.actionName
	b.Layout = "admin/layout.html"
	b.o = orm.NewOrm()
	b.TplName = b.viewsTplName
	if b.rouleName == "login/index" || b.rouleName == "login/logout" {
		if b.GetSession("admin_user_go") != nil {
			b.History("已登录系统", "/admin/index")
		}
	} else {
		/*
			if b.GetSession("admin_user_go") == nil {
				b.History("未登录", "/admin/login/index")
			} else {
				userSess := b.GetSession("admin_user_go")
				userJson, _ := json.Marshal(userSess)
				_ = json.Unmarshal(userJson, &b.loginUser)

				if b.rouleName == "admin/index" || b.rouleName == "admin/menu" {
					//菜单权限
					b.MenuRole()
				} else {
					//菜单权限
					b.MenuRole()
					//菜单权限判断
					//logs.Error("menu_id====", b.showMenuId)
					b.Roule()
				}

				//logs.Error("user_id====", userJson)
				//logs.Error("user_id====", b.loginUser.Id)
			}
		*/
	}
	//logs.Error("rouleName====", b.rouleName)

}
func (b *baseController) MenuRole() {

	if b.loginUser.Level != 1 {
		groupId := b.loginUser.GroupId
		groupS := strings.Split(groupId, ",")
		if groupId != "" {
			var listH orm.ParamsList
			menuIdS := "0"
			num, err := b.o.QueryTable(new(models.Group).TableName()).Filter("id__in", groupS).Filter("status__in", 1).ValuesFlat(&listH, "Access")
			if err == nil {
				if num > 0 {
					for _, value := range listH {
						menuIdS = menuIdS + "," + value.(string)
					}
				}
				//logs.Error("menuIdS", menuIdS)

				menuS := strings.Split(menuIdS, ",")
				num, _ := b.o.QueryTable(new(models.Menu).TableName()).Filter("id__in", menuS).Filter("status__in", 1).ValuesFlat(&listH, "Path")
				if num > 0 {
					for _, value := range listH {
						menuIdS = menuIdS + "," + value.(string)
					}
				}

				b.showMenuId = menuIdS

			}

		}
	} else {

		var listHa orm.ParamsList
		menuIdS := "0"
		num, _ := b.o.QueryTable(new(models.Menu).TableName()).Filter("status__in", 1).ValuesFlat(&listHa, "Id")
		if num > 0 {
			for _, valuet := range listHa {
				//logs.Error("LEVELsssss====", reflect.TypeOf(valuet))
				menuIdS = menuIdS + "," + strconv.Itoa(int(valuet.(int64)))
			}

		}

		b.showMenuId = menuIdS
	}

}
func (b *baseController) Roule() {
	cate := models.Menu{Url: b.rouleName}
	b.o.Read(&cate, "Url")
	if cate.Id > 0 {
		menuS := strings.Split(b.showMenuId, ",")
		if !common.InArray(menuS, strconv.Itoa(cate.Id)) {
			b.Erro("无些菜单权限", "权限管理", 0, b.resData)
		}
	} else {
		//b.Erro("无权限信息", "权限管理", 0)
	}
}

//topjui  操作成功返回
func (b *baseController) Succ(msg string, title string, data map[string]string) {

	re := make(map[string]string)
	if msg == "" {
		re["message"] = "恭喜你，操作成功！"
	} else {
		re["message"] = msg
	}
	if title == "" {
		re["title"] = "操作提示"
	} else {
		re["title"] = title
	}

	/*使用键输出地图值 */
	for country := range data {
		re[country] = data[country]
	}

	re["statusCode"] = "200"
	b.Data["json"] = re
	b.ServeJSON()
}

//topjui  操作失败返回
func (b *baseController) Erro(msg string, title string, code int, data map[string]string) {
	re := make(map[string]string)
	if msg == "" {
		re["message"] = "恭喜你，操作成功！"
	} else {
		re["message"] = msg
	}
	if title == "" {
		re["title"] = "操作提示"
	} else {
		re["title"] = title
	}

	/*使用键输出地图值 */
	for country := range data {
		re[country] = data[country]
	}

	if code == 0 {
		re["statusCode"] = "404"
	} else {
		re["statusCode"] = strconv.Itoa(code)
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

//获取用户IP地址
func (b *baseController) getClientIp() string {
	//s := strings.Split(b.Ctx.Request.RemoteAddr, ":")
	//return s[0]
	return b.Ctx.Request.RemoteAddr
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

/*
md5加密
*/
func MyMd5(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}
