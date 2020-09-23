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

func (this *ProductController) Index() {
	ProductList, err := ProductModel.GetProductList(0, 10)
	if err != nil {
		this.Error(err)
		return
	}
	this.Data['ProductList'] = ProductList
	fmt.Println(ProductList)
	this.TplName = "newSecKill2/product/index.html"
	return
}
