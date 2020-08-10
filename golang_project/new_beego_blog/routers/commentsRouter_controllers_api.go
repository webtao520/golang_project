package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)
// 提供接口的几种常用方式
func init() {
   // 注册路由，设置请求方式，设置处理路由函数
	beego.GlobalControllerRouter["new_beego_blog/controllers/api:FrontController"] = append(beego.GlobalControllerRouter["new_beego_blog/controllers/api:FrontController"],
		beego.ControllerComments{
			Method:           "GetCaptcha",
			Router:           `/captcha`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["new_beego_blog/controllers/api:FrontController"] = append(beego.GlobalControllerRouter["new_beego_blog/controllers/api:FrontController"],
		beego.ControllerComments{
			Method:           "PostCaptchaCheck",
			Router:           `/captcha/check`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["new_beego_blog/controllers/api:FrontController"] = append(beego.GlobalControllerRouter["new_beego_blog/controllers/api:FrontController"],
		beego.ControllerComments{
			Method:           "GetTest",
			Router:           `/test/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
