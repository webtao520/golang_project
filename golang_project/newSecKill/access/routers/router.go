package routers

import (
	"newSecKill/access/controllers/index"
	"newSecKill/access/controllers/seckill"
	"newSecKill/access/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/", &index.IndexController{}, "*:Index")
	beego.Router("/index/index", &index.IndexController{}, "*:Index")
	beego.Router("/index/register", &user.UserController{}, "*:Register")
	beego.Router("/seckill/index", &seckill.SecKillController{}, "*:Index")
}
