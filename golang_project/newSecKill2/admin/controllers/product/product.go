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

func (this *ProductController) Success(url string, msg interface{}) {
	this.Data["success"] = fmt.Sprintln(msg)
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
	this.TplName = "product/index.html"
	return
}

// 添加产品
func (this *ProductController) AddProduct() {
	if this.Ctx.Input.IsPost() {
		ProductName := this.GetString("ProductName")
		if (len(ProductName)) <= 0 {
			err := "ProductName 不能为空"
			this.Error(err)
			return
		}
		num, _, err := ProductModel.GetProductListByName(ProductName, 1)
		if err != nil {
			this.Error(err)
			return
		}
		if num == 1 {
			this.Error("您输入的产品已存在")
			return
		}
		Total, err := this.GetInt("Total")
		if err != nil || Total < 0 {
			this.Error("Total 必须是大于0的整数")
			return
		}
		Status, err := this.GetInt("Status")
		if err != nil {
			this.Error(fmt.Sprintln("Status  err : %v", err))
			return
		}
		product := models.SecKillProduct{
			ProductName: ProductName,
			Total:       Total,
			Status:      Status,
		}
		id, err := ProductModel.InsertProduct(&product)
		if err != nil {
			this.Error(err)
			return
		} else {
			fmt.Println(id)
		}
		this.Success("index", "添加成功")
	} else {
		this.TplName = "product/add.html"
	}
}
