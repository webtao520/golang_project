package admin

import "fmt"

// 定义首页控制器
type IndexController struct {
	baseController
}

// 后台首页
func (this *IndexController) Index(){
	fmt.Println("后台首页.....")
    this.display()
}