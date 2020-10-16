package admin

import (
	"beego_blog/util"
	"beego_weihuaijing/models"
	"github.com/astaxie/beego/logs"
	"os"
	"strings"
	"time"
)

type AdminController struct {
	baseController
}
type Menu struct {
	Id      int    `json:"id"`
	Pid     int    `json:"pid"`
	Text    string `json:"text"`
	State   string `json:"state"`
	IconCls string `json:"iconCls"`
	Url     string `json:"url"`
}

func (a *AdminController) Index() {
	//logs.Error("views====", a.viewsTplName)
	//menuS := strings.Split(a.showMenuId, ",")
	var list []models.Menu
	//a.o.QueryTable(new(models.Menu).TableName()).Filter("id__in", menuS).Filter("status__in", 1).Filter("pid", 0).OrderBy("-MenuSort").All(&list)
	a.o.QueryTable(new(models.Menu).TableName()).Filter("status__in", 1).Filter("pid", 0).OrderBy("-MenuSort").All(&list)
	a.Data["lists"] = list
	a.Layout = ""
	a.TplName = "admin/index.html"
}

func (a *AdminController) Menu() {
	menuId, _ := a.GetInt("menu")
	//menuS := strings.Split(a.showMenuId, ",")
	//levelId, _ := a.GetInt("levelId")
	var list []models.Menu
	//a.o.QueryTable(new(models.Menu).TableName()).Filter("id__in", menuS).Filter("status__in", 1).Filter("pid", menuId).OrderBy("-MenuSort").All(&list)
	a.o.QueryTable(new(models.Menu).TableName()).Filter("status__in", 1).Filter("pid", menuId).OrderBy("-MenuSort").All(&list)

	a.Data["json"] = list
	a.ServeJSON()
}

//上传接口
func (c *AdminController) Upload() {

	f, h, err := c.GetFile("file")
	var bs = 0
	img := ""
	re := make(map[string]string)
	if err == nil {
		exStrArr := strings.Split(h.Filename, ".")
		exStr := strings.ToLower(exStrArr[len(exStrArr)-1])
		if exStr != "jpg" && exStr != "png" && exStr != "gif" {
			bs = 1
		}
		//创建目录
		uploadDir := "static/upload/" + time.Now().Format("2006/01/02/")
		err := os.MkdirAll(uploadDir, 0777)
		if err != nil {

			logs.Error("=====", err)
		}
		img = uploadDir + util.UniqueId() + "." + exStr
		c.SaveToFile("file", img) // 保存位置在 static/upload, 没有文件夹要先创建
		//logs.Error("img======", img)
		re["filePath"] = "/" + img
		bs = 0
	} else {
		bs = 2
	}
	defer f.Close()

	if bs == 0 {
		c.Succ("上传成功", "上传提示", re)
	} else if bs == 1 {
		c.Erro("上传只能.jpg 或者png格式", "上传提示", 0, c.resData)
	} else {
		c.Erro("上传异常"+err.Error(), "上传提示", 0, c.resData)
	}

}
