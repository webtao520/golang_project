package admin

import (
	//"fmt"
	"strings"

	"github.com/astaxie/beego"
)

type baseController struct {
	beego.Controller // 加载beego核心控制器
	userid           int64
	username         string
	permissionlist   map[string]int
	moduleName       string
	controllerName   string
	actionName       string
}

// 这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些 Method 方法之前执行，用户可以重写这个函数实现类似用户验证之类。
func (this *baseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.moduleName = "admin"
	//fmt.Printf("%T",controllerName)
	//fmt.Println(controllerName[0 : len(controllerName)-10])
	//this.controllerName=strings.ToLower()
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
}

// 后台模板渲染
func (this *baseController) display(tpl ...string) {
	var tplname string
	if len(tpl) == 1 {
		tplname = this.moduleName + "/" + tpl[0] + ".html"
	} else {
		tplname = this.moduleName + "/" + this.controllerName + "_" + this.actionName + ".html"
	}
	this.Layout = this.moduleName + "/layout.html"
	this.Data["version"] = beego.AppConfig.String("AppVer")
	this.Data["adminid"] = this.userid
	this.Data["adminname"] = this.username
	this.Data["userpermission"] = this.permissionlist
	this.TplName = tplname

}

// 创建json结构体数据
type RetJson struct {
	Status bool        `json:"status"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

func RetResource(status bool, data interface{}, msg string) (apijson *RetJson) {
	apijson = &RetJson{Status: status, Data: data, Msg: msg}
	return
}
