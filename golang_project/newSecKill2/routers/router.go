package routers

import (
	"github.com/beego/admin"
)

func init() {
	admin.Run()
	//beego.Router("/", &controllers.MainController{})
}
