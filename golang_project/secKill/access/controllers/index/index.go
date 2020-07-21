package index

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Index() {
	user := this.GetSession("user")
	this.Data["user"] = user
	if user != nil {
		this.Data["Login"] = true
	}
	this.TplName = "index/index.html"
	return
}
