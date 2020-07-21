package routers

import (
	"github.com/astaxie/beego"
	"github.com/beego/admin"
)

func init() {
	admin.Run()
	beego.Router("/")
	//beego.Router("/", &controllers.MainController{})
}
