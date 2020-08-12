package routers

import (
	"new_beego_blog/controllers/admin"
	"new_beego_blog/controllers/api"
	"new_beego_blog/controllers/blog"

	"github.com/astaxie/beego"
)

func init() {
	//前台路由
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/", &blog.MainController{}, "*:Index")

	// 博客后台路由
	beego.Router("/admin", &admin.IndexController{}, "*:Index")
	beego.Router("/admin/login", &admin.AccountController{}, "post:Login")
	beego.Router("/admin/register", &admin.AccountController{}, "post:Register")

	// 使用命名空间

	//front api
	front := beego.NewNamespace("/v1",
		beego.NSNamespace("/front",
			beego.NSInclude(
				&api.FrontController{},
			),
		),
	)
	//login api
	login := beego.NewNamespace("/v1",
		beego.NSNamespace("/login",
			beego.NSInclude(
				&api.FrontController{},
			),
		),
	)
	beego.AddNamespace(front, login)
}
