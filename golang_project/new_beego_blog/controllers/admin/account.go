package admin

import (
	"fmt"
	"new_beego_blog/models"
)

type AccountController struct {
	baseController
}

// 注册
func (this *AccountController) Register() {
	aul:=new(models.UserJson) // 分配内存
	fmt.Println(aul)
} 
