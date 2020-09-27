package routers

import (
	"newSecKill2/admin/controllers/activity"
	"newSecKill2/admin/controllers/product"

	"github.com/astaxie/beego"
	"github.com/beego/admin"
)

func init() {
	admin.Run()
	// 产品
	beego.Router("/seckill/product/index", &product.ProductController{}, "*:Index")                 // 产品列表页
	beego.Router("/seckill/product/addProuct", &product.ProductController{}, "*:AddProduct")        // 添加产品
	beego.Router("/seckill/product/updateProduct", &product.ProductController{}, "*:UpdateProduct") //更新产品
	beego.Router("/seckill/product/delProduct", &product.ProductController{}, "*:DelProduct")       // 删除商品

	// 活动
	beego.Router("/seckill/activity/index", &activity.ActivityController{}, "*:Index")                   // 活动列表
	beego.Router("/seckill/activity/addActivity", &activity.ActivityController{}, "*:AddActivity")       // 添加活动
	beego.Router("/seckill/activity/updateActivity", &activity.ActivityController{}, "*:UpdateActivity") // 修改活动
	beego.Router("/seckill/activity/delActivity", &activity.ActivityController{}, "*:DelActivity")       // 删除活动
}
