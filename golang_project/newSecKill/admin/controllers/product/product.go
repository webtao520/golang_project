package product

import (
	"fmt"
	"newSecKill/admin/models"

	"github.com/astaxie/beego"
)

var (
	ProductModel *models.SecKillProduct = models.NewProductModel()
)

// 定义产品控制器
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
	// 列表数据
	ProductList, err := ProductModel.GetProductList(0, 10)
	if err != nil {
		this.Error(err)
		return
	}
	this.Data["ProductList"] = ProductList
	this.TplName = "seckill/product/index.html"
	return
}

// 添加产品
func (this *ProductController) AddProduct() {
	if this.Ctx.Input.IsPost() {
		ProductName := this.GetString("ProductName")
		if len(ProductName) <= 0 {
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
			this.Error("Total 必须大于0的整数")
			return
		}

		Status, err := this.GetInt("Status")
		if err != nil {
			this.Error(fmt.Sprintln("Status err : %v", err))
			return
		}

		//fmt.Println("=============",ProductModel)
		produnct := models.SecKillProduct{
			ProductName: ProductName,
			Total:       Total,
			Status:      Status,
		}

		// fmt.Println("==============",produnct) //============== {0 iPhone X 23 1}
		// 保存产品数据
		id, err := ProductModel.InsertProduct(&produnct)
		if err != nil {
			this.Error(err)
			return
		}
		fmt.Println("插入id====>", id)
		this.Success("index", "添加成功")
	} else {
		this.TplName = "seckill/product/add.html"
		return
	}
}

// 更新产品
func (this *ProductController) UpdateProduct() {

}

// 删除产品
func (this *ProductController) DelProduct() {

}
