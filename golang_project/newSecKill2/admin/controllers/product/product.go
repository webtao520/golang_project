package product

import (
	"fmt"
	"newSecKill2/admin/models"

	"github.com/astaxie/beego"
)

var ProductModel *models.SecKillProduct = models.NewProductModel()

type ProductController struct {
	beego.Controller
}

func (this *ProductController) Error(err interface{}) {
	// 获取当前页面url
	url := this.Ctx.Request.Referer()
	this.Data["error"] = fmt.Sprintln(err)
	this.Redirect(url, 302)
	return
}

func (this *ProductController) Index() {
	ProductList, err := ProductModel.GetProductList(0, 10)
	if err != nil {
		this.Error(err)
		return
	}
	this.Data["ProductList"] = ProductList
	this.TplName = "newSecKill2/product/index.html"
	return
}
