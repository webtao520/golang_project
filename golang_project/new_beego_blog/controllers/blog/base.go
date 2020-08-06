package blog

import (
	//"fmt"

	//"fmt"

	"github.com/astaxie/beego"
)

type baseController struct {
	beego.Controller
	options      map[string]string
	right string
}

func (this *baseController) Prepare(){
	this.Data["hidejs"] = `<!--[if lt IE 9]>
	<script src="/static/js/html5shiv.min.js"></script>
	<![endif]-->`
}

func (this *baseController) display(tpl string){
	theme := "double"
	// 类型断言
	if v, ok := this.options["theme"]; ok && v != "" {
		theme = v
	}
	this.Layout=theme + "/layout.html"
	this.Data["root"]="/"+beego.BConfig.WebConfig.ViewsPath + "/" + theme + "/"
	this.TplName=theme+"/" + tpl +".html"
	this.LayoutSections=make(map[string]string)
	this.LayoutSections["head"]=theme+"/head.html"
	if tpl == "index" {
		this.LayoutSections["banner"] = theme + "/banner.html"
	}
	if this.right != "" {
		this.LayoutSections["right"]=theme + "/" +this.right
	}
	this.LayoutSections["foot"]=theme+"/foot.html"
}