package index

import (
	"fmt"

	"github.com/astaxie/beego"
)

// 定义 access 首页控制器
type IndexController struct {
	beego.Controller
}

func (this *IndexController) Index() {
	fmt.Println("秒杀首页.....")
	this.TplName = "index/index.html"
	return
}
