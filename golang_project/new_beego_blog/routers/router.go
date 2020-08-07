package routers

import (
	"new_beego_blog/controllers/admin"
	"new_beego_blog/controllers/blog"

	"github.com/astaxie/beego"
)

func init() {
	//前台路由
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/", &blog.MainController{}, "*:Index")

	// 博客后台路由
	beego.Router("/admin", &admin.IndexController{}, "*:Index")
	beego.Router("/admin/register", &admin.AccountController{}, "post:Register")
}
