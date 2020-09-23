package routers

import (
	"newSecKill2/admin/controllers/product"

	"github.com/astaxie/beego"
	"github.com/beego/admin"
)

func init() {
	admin.Run()
	// 产品
	beego.Router("/newSecKill2/product/index", &product.ProductController{}, "*:Index") // 产品列表页
	//beego.Router()                                                                  // 添加产品
}
