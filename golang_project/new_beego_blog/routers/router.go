package routers

import (
	"new_beego_blog/controllers/blog"

	"github.com/astaxie/beego"
)

func init() {
	//前台路由
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/", &blog.MainController{}, "*:Index")
}
