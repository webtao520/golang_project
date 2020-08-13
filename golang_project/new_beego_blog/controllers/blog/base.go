package blog

import (
	//"fmt"
	"new_beego_blog/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type baseController struct {
	beego.Controller
	options  map[string]string
	right    string
	page     int
	pagesize int
}

//  在所有方法之前加载
func (this *baseController) Prepare() {
	this.Data["IsLogin"] = this.IsLogin()
	this.options = models.GetOptions()
	this.Data["hidejs"] = `<!--[if lt IE 9]>
	<script src="/static/js/html5shiv.min.js"></script>
	<![endif]-->`

	var (
		pagesize int
		err      error
		page     int
	)
	if page, err = strconv.Atoi(this.Ctx.Input.Param(":page")); err != nil || page < 1 {
		page = 1
	}

	if pagesize, err = strconv.Atoi(this.getOption("pagesize")); err != nil || pagesize < 1 {
		pagesize = 10
	}
	this.page = page
	this.pagesize = pagesize
}

// 判断是否登陆
func (this *baseController) IsLogin() bool {
	//session预先判断是否登录
	userId := this.GetSession("userId")
	if userId != nil && userId.(int64) > 0 {
		return true
	} else {
		arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
		if len(arr) == 2 {
			idstr, password := arr[0], arr[1]
			userid, _ := strconv.ParseInt(idstr, 10, 0)
			if userid > 0 {
				var user models.User
				user.Id = userid
				if user.Read() == nil && password == models.Md5([]byte(this.getClientIp()+"|"+user.Password)) {
					this.SetSession("userId", user.Id)
					this.SetSession("userName", user.Username)
					return true
				}
			}
		}
	}
	return false
}

//获取用户IP地址
func (this *baseController) getClientIp() string {
	s := this.Ctx.Request.Header.Get("X-Real-IP")
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

func (this *baseController) getOption(name string) string {
	if v, ok := this.options[name]; ok {
		return v
	} else {
		return ""
	}
}

func (this *baseController) setHeadMetas(params ...string) {
	title_buf := make([]string, 0, 3)
	if len(params) == 0 && this.getOption("sitename") != "" {
		title_buf = append(title_buf, this.getOption("sitename"))
	}
	if len(params) > 0 {
		title_buf = append(title_buf, params[0])
	}
	title_buf = append(title_buf, this.getOption("subtitle"))
	this.Data["title"] = strings.Join(title_buf, " - ")

	if len(params) > 1 {
		this.Data["keywords"] = params[1]
	} else {
		this.Data["keywords"] = this.getOption("keywords")
	}

	if len(params) > 2 {
		this.Data["description"] = params[2]
	} else {
		this.Data["description"] = this.getOption("description")
	}
}

/*
	对于一个复杂的 LayoutContent，其中可能包括有javascript脚本、CSS 引用等，根据惯例，
	通常 css 会放到 Head 元素中，javascript 脚本需要放到 body 元素的末尾，而其它内容则根据需要放在合适的位置。
	在 Layout 页中仅有一个 LayoutContent 是不够的。所以在 Controller 中增加了一个 LayoutSections属性，可以允许 Layout 页中设置多个 section，
	然后每个 section 可以分别包含各自的子模板页。
*/

func (this *baseController) display(tpl string) {
	theme := "double"
	if v, ok := this.options["theme"]; ok && v != "" {
		theme = v
	}
	this.Layout = theme + "/layout.html"
	this.Data["root"] = "/" + beego.BConfig.WebConfig.ViewsPath + "/" + theme + "/" // /views/double/
	this.TplName = theme + "/" + tpl + ".html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["head"] = theme + "/head.html"
	if tpl == "index" {
		this.LayoutSections["banner"] = theme + "/banner.html"
	}
	if this.right != "" {
		this.LayoutSections["right"] = theme + "/" + this.right
	}
	this.LayoutSections["foot"] = theme + "/foot.html"
}
