package admin

// 用户管理控制器
type UserController struct {
	baseController
}

// 添加用户
func (this *UserController) Add(){
	this.display()
}
