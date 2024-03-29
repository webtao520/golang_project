package admin

import (
	"new_beego_blog/models"
	"os"
	"runtime"
)

// 定义首页控制器
type IndexController struct {
	baseController
}

// 后台首页
func (this *IndexController) Index() {
	this.Data["hostname"], _ = os.Hostname()
	this.Data["gover"] = runtime.Version()
	this.Data["os"] = runtime.GOOS
	this.Data["cpunum"] = runtime.NumCPU()
	this.Data["arch"] = runtime.GOARCH
	this.Data["postnum"], _ = new(models.Post).Query().Count()
	this.Data["tagnum"], _ = new(models.Tag).Query().Count()
	this.Data["usernum"], _ = new(models.User).Query().Count()
	this.display()
}
