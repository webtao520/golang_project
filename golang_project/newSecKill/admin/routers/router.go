package routers

import (
	"newSecKill/admin/controllers/activity"
	"newSecKill/admin/controllers/product"
	_ "newSecKill/admin/models"

	"github.com/astaxie/beego"
	"github.com/beego/admin"
)

func init() {
	admin.Run()
	//产品
	beego.Router("/seckill/product/index", &product.ProductController{}, "*:Index")                 // 列表
	beego.Router("/seckill/product/addProuct", &product.ProductController{}, "*:AddProduct")        // 增加
	beego.Router("/seckill/product/updateProduct", &product.ProductController{}, "*:UpdateProduct") // 更新
	beego.Router("/seckill/product/delProduct", &product.ProductController{}, "*:DelProduct")       // 删除

	// 活动
	beego.Router("/seckill/activity/index", &activity.ActivityController{}, "*:Index")
	beego.Router("/seckill/activity/addActivity", &activity.ActivityController{}, "*:AddActivity")       // 添加活动
	beego.Router("/seckill/activity/updateActivity", &activity.ActivityController{}, "*:UpdateActivity") // 修改活动
	beego.Router("/seckill/activity/delActivity", &activity.ActivityController{}, "*:DelActivity")       // 删除活动
}
