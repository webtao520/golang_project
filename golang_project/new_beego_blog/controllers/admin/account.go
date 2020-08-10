package admin

import (
	"encoding/json"
	"new_beego_blog/models"

	"github.com/astaxie/beego/logs"
)

type AccountController struct {
	baseController
}

// 注册 beego api 开发
func (this *AccountController) Register() {
	aul := new(models.UserJson) // 分配内存
	//  在 API 的开发中，我们经常会用到 JSON 或 XML 来作为数据交互的格式，
	//如何在 beego 中获取 Request Body 里的 JSON 或 XML 的数据呢？
	data := this.Ctx.Input.RequestBody
	//fmt.Println("==================", this.Ctx.Input.RequestBody)
	if err := json.Unmarshal(data, &aul); err != nil {
		logs.Warning(err.Error())
		this.Data["json"] = RetResource(false, nil, "参数错误")
		this.ServeJSON()
		return
	}

}
