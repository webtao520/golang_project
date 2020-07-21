package routers

import (
	"secKill/access/controllers/index"
	"secKill/access/controllers/seckill"
	"secKill/access/controllers/user"
	_ "secKill/access/models"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &index.IndexController{}, "*:Index")
	beego.Router("/index/index", &index.IndexController{}, "*:Index")
	beego.Router("/index/register", &user.UserController{}, "*:Register")

	beego.Router("/seckill/user/login", &user.UserController{}, "*:Login")
	beego.Router("/seckill/user/exit", &user.UserController{}, "*:Exit")

	beego.Router("/seckill/index", &seckill.SecKillController{}, "*:Index")
	beego.Router("/seckill/seckill/:ActivityId", &seckill.SecKillController{}, "*:SecKill")

	// 验证用户是否已经登录
	beego.InsertFilter("/seckill/*", beego.BeforeExec, FilterSecKill)
}

var FilterSecKill = func(ctx *context.Context) {
	user := ctx.Input.Session("user")
	if user == nil && ctx.Request.RequestURI != "/seckill/user/login" {
		ctx.Redirect(302, "/seckill/user/login")
	}
}
