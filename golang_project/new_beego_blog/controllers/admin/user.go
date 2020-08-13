package admin

import "strings"

// 用户管理控制器
type UserController struct {
	baseController
}

// 添加用户
func (this *UserController) Add() {
	input := make(map[string]string)
	errmsg := make(map[string]string)
	if this.Ctx.Request.Method == "POST" {
		username := strings.TrimSpace()
	}

	this.display()
}
