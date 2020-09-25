package routers

import (
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
}
