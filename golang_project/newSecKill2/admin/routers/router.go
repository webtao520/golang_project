package routers

import (
	"newSecKill2/admin/controllers/product"

	"github.com/astaxie/beego"
	"github.com/beego/admin"
)

func init() {
	admin.Run()
	// 产品
	beego.Router("/seckill/product/index", &product.ProductController{}, "*:Index")
}
