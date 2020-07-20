package routers

import (
	"secKill/admin/controllers/activity"
	"secKill/admin/controllers/product"
	_ "secKill/admin/models"

	"github.com/astaxie/beego"
	"github.com/beego/admin"
)

func init() {
	admin.Run()
	// beego.Router("/", &controllers.MainController{})
	// 产品
	beego.Router("/seckill/product/index", &product.ProductController{}, "*:Index")
	beego.Router("/seckill/product/addProuct", &product.ProductController{}, "*:AddProduct")
	beego.Router("/seckill/product/updateProduct", &product.ProductController{}, "*:UpdateProduct")
	beego.Router("/seckill/product/delProduct", &product.ProductController{}, "*:DelProduct")

	// 活动
	beego.Router("/seckill/activity/index", &activity.ActivityController{}, "*:Index")
	beego.Router("/seckill/activity/addActivity", &activity.ActivityController{}, "*:AddActivity")
	beego.Router("/seckill/activity/updateActivity", &activity.ActivityController{}, "*:UpdateActivity")
	beego.Router("/seckill/activity/delActivity", &activity.ActivityController{}, "*:DelActivity")
}
