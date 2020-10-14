package routers

import (
	"newSecKill/access/controllers/index"
	"newSecKill/access/controllers/seckill"
	"newSecKill/access/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &index.IndexController{}, "*:Index")
	beego.Router("/index/index", &index.IndexController{}, "*:Index")
	beego.Router("/index/register", &user.UserController{}, "*:Register")
	beego.Router("/seckill/user/login", &user.UserController{}, "*:Login")

	beego.Router("/seckill/index", &seckill.SecKillController{}, "*:Index")                //秒杀活动列表页面
	beego.Router("seckill/seckill/:ActivityId", &seckill.SecKillController{}, "*:SecKill") // 秒杀
}
