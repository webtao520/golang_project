package main

import (
	_ "new_beego_blog/models"
	_ "new_beego_blog/routers"

	"github.com/astaxie/beego"
)

//添加自定义模板函数
func hasPermission(permissionlist map[string]int, value int) (out bool) {
	for _, id := range permissionlist {
		if value == id {
			return true
		}
	}
	return false
}

func main() {
	beego.AddFuncMap("haspermission", hasPermission) // beego 支持用户定义模板函数，但是必须在 beego.Run() 调用之前
	beego.Run()
}
