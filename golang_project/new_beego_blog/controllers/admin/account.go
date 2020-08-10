package admin

import (
	"encoding/json"
	"fmt"
	"new_beego_blog/models"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
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
	code := this.GetSession("Captcha")
	if code == nil {
		logs.Warning("Session中没有获取到验证码")
		this.Data["json"] = RetResource(false, nil, "请求异常")
		this.ServeJSON()
		return
	}

	if code != aul.Captcha {
		logs.Warning(fmt.Sprintf("sessCaptcha:%s,inputCaptcha:%s", code, aul.Captcha))
		this.Data["json"] = RetResource(false, nil, "验证码不正确")
		this.ServeJSON()
		return
	}
	valid := validation.Validation{}
	if v := valid.Required(aul.Username1, "username"); !v.Ok {
		this.Data["json"] = RetResource(false, nil, "请输入用户名")
		this.ServeJSON()
		return
	}
	if v := valid.MaxSize(aul.Username1, 30, "username"); !v.Ok {
		this.Data["json"] = RetResource(false, nil, "用户名长度不能大于15个字符")
		this.ServeJSON()
		return
	}
	if !checkUsername(aul.Username1) {
		this.Data["json"] = RetResource(false, nil, "输入的用户名不符合要求(仅允许字母开头,并以字母、数字、-、_组成)")
		this.ServeJSON()
		return
	}
}

func checkUsername(username string) (b bool) {

}
