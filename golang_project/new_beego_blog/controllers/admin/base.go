package admin

import (
	//"fmt"
	"fmt"
	"new_beego_blog/models"
	"strconv"
	"strings"
	"time"

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
	this.auth() // 登陆状态验证 显示菜单项
	//this.checkPermission()
}

//登录状态验证
func (this *baseController) auth() {
	//允许任何人默认拥有访问account，comments的权限
	this.permissionlist = map[string]int{"account": 0, "comments": 0}
	//session预先判断是否登录
	if userId := this.GetSession("userId"); userId != nil {
		this.userid = userId.(int64)
	}
	if this.userid > 0 {
		if userName := this.GetSession("userName"); userName != nil {
			this.username = userName.(string)
		}
		var user models.User
		user.Id = this.userid
		if user.Read() == nil {
			this.permissionlist = MakePermissionList(user)
		}
	} else {
		//未登录时，将获取cookie，如cookie存在则认证1次并保存相关信息到session，
		arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
		if len(arr) == 2 {
			idstr, password := arr[0], arr[1]
			userid, _ := strconv.ParseInt(idstr, 10, 0)
			if userid > 0 {
				var user models.User
				user.Id = userid
				if user.Read() == nil && password == models.Md5([]byte(this.getClientIp()+"|"+user.Password)) {
					this.userid = user.Id
					this.username = user.Username
					this.SetSession("userId", user.Id)
					this.SetSession("userName", user.Username)
					this.permissionlist = MakePermissionList(user)
				}
			}
		}
	}
	if this.userid > 0 && this.actionName == "login" {
		this.Redirect("/admin", 302)
	}
}

//渲染模版
func (this *baseController) display(tpl ...string) {
	var tplname string
	if len(tpl) == 1 {
		tplname = this.moduleName + "/" + tpl[0] + ".html"
	} else {
		tplname = this.moduleName + "/" + this.controllerName + "_" + this.actionName + ".html"
	}
	this.Data["version"] = beego.AppConfig.String("AppVer")
	this.Data["adminid"] = this.userid
	this.Data["adminname"] = this.username
	this.Data["userpermission"] = this.permissionlist
	fmt.Println("后台管理菜单数据信息", this.userid, this.username, this.permissionlist)
	this.Layout = this.moduleName + "/layout.html"
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

// 获取用ip地址
func (this *baseController) getClientIp() string {
	s := this.Ctx.Request.Header.Get("X-Real-IP")
	fmt.Println("====X-Real-IP=====", s)
	if s == "" {
		forwarded := this.Ctx.Request.Header.Get("X-Forwarded-For")
		if forwarded != "" {
			list := strings.Split(forwarded, ":")
			if len(list) > 0 {
				s = list[0]
			}
		} else {
			s = strings.Split(this.Ctx.Request.RemoteAddr, ":")[0]
		}
	}
	return s
}

func (this *baseController) getTime() time.Time {
	options := models.GetOptions()
	timezone := float64(0)
	if v, ok := options["timezone"]; ok {
		timezone, _ = strconv.ParseFloat(v, 64)
	}
	add := timezone * float64(time.Hour)
	return time.Now().UTC().Add(time.Duration(add))
}

// 显示错误提示
func (this *baseController) showmsg(msg ...string) {
	redirect := this.Ctx.Request.Referer()
	if redirect == "" {
		redirect = "/admin"
	}
	if len(msg) == 1 {
		if this.userid == 0 {
			redirect = "/"
		}
		msg = append(msg, redirect)
	}
	this.Data["adminid"] = this.userid
	this.Data["adminname"] = this.username
	this.Data["msg"] = msg[0]
	this.Data["redirect"] = msg[1]
	this.Layout = this.moduleName + "/layout.html"
	this.TplName = this.moduleName + "/" + "showmsg.html"
	_ = this.Render()
	this.StopRun()
}

func (this *baseController) Isdefaultsrc(value string) bool {
	var defaultdir = "/static/upload/default/"
	if value != "" {
		if index := strings.Index(value, defaultdir); index != -1 {
			return true
		}
	}
	return false
}
