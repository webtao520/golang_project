package admin

import (
	"encoding/json"
	"fmt"
	"new_beego_blog/models"
	"regexp"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
)

type AccountController struct {
	baseController
}

// 登录
func (this *AccountController) Login() {
	aul := new(models.UserJson)
	data := this.Ctx.Input.RequestBody
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
	if aul.Username == "" {
		this.Data["json"] = RetResource(false, nil, "请输入帐号")
		this.ServeJSON()
		return
	}
	if aul.Password == "" {
		this.Data["json"] = RetResource(false, nil, "请输入密码")
		this.ServeJSON()
		return
	}
	var user models.User
	user.Username = aul.Username
	if user.Read("username") != nil || user.Password != models.Md5([]byte(aul.Password)) {
		this.Data["json"] = RetResource(false, nil, "帐号或密码错误")
		this.ServeJSON()
		return
	}
	if user.Active == 0 {
		this.Data["json"] = RetResource(false, nil, "该帐号未激活")
		this.ServeJSON()
		return
	}
	user.Logincount += 1
	user.Lastip = this.getClientIp()
	user.Lastlogin = this.getTime()
	_ = user.Update()
	authkey := models.Md5([]byte(this.getClientIp() + "|" + user.Password))
	this.Ctx.SetCookie("auth", strconv.FormatInt(user.Id, 10)+"|"+authkey, 86400)
	this.SetSession("userId", user.Id)
	this.SetSession("userName", user.Username)
	logs.Info(fmt.Sprintf("userid:%d,username:%s,登录成功", user.Id, user.Username))
	this.Data["json"] = RetResource(true, nil, "登录成功")
	this.ServeJSON()
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
	if v := valid.Required(aul.Nickname, "nickname"); !v.Ok {
		this.Data["json"] = RetResource(false, nil, "请输入昵称")
		this.ServeJSON()
		return
	}
	if v := valid.MaxSize(aul.Nickname, 30, "nickname"); !v.Ok {
		this.Data["json"] = RetResource(false, nil, "昵称长度不能大于15个字符")
		this.ServeJSON()
		return
	}
	if v := valid.Required(aul.Password1, "password"); !v.Ok {
		this.Data["json"] = RetResource(false, nil, "请输入密码")
		this.ServeJSON()
		return
	}
	if !checkPassword(aul.Password1) {
		this.Data["json"] = RetResource(false, nil, "输入的密码不符合要求(仅允许字母、数字和部分符号组成)")
		this.ServeJSON()
		return
	}
	if v := valid.Required(aul.Password2, "password2"); !v.Ok {
		this.Data["json"] = RetResource(false, nil, "请再次输入密码")
		this.ServeJSON()
		return
	}
	if aul.Password1 != aul.Password2 {
		this.Data["json"] = RetResource(false, nil, "两次输入的密码不一致")
		this.ServeJSON()
		return
	}
	if !checkPassword(aul.Password2) {
		this.Data["json"] = RetResource(false, nil, "输入的密码不符合要求(仅允许字母、数字和部分符号组成)")
		this.ServeJSON()
		return
	}
	if v := valid.Required(aul.Email, "email"); !v.Ok {
		this.Data["json"] = RetResource(false, nil, "请输入email地址")
		this.ServeJSON()
		return
	}
	if v := valid.Email(aul.Email, "email"); !v.Ok {
		this.Data["json"] = RetResource(false, nil, "Email无效")
		this.ServeJSON()
		return
	}

	var user models.User
	err := user.Query().Filter("username", aul.Username1).One(&user)
	if err == nil {
		logs.Warning(fmt.Sprintf("用户名:%s 已被注册", aul.Username1))
		this.Data["json"] = RetResource(false, nil, fmt.Sprintf("用户名:%s 已被注册", aul.Username1))
		this.ServeJSON()
		return
	}
	err = user.Query().Filter("nickname", aul.Nickname).One(&user)
	if err == nil {
		logs.Warning(fmt.Sprintf("昵称:%s 已被使用", aul.Nickname))
		this.Data["json"] = RetResource(false, nil, fmt.Sprintf("昵称:%s 已被使用", aul.Nickname))
		this.ServeJSON()
		return
	}
	user.Username = aul.Username1
	user.Password = models.Md5([]byte(aul.Password1))
	user.Email = aul.Email
	user.Active = int8(1)
	user.Upcount = int64(3)
	user.Lastip = this.getClientIp() // 获取请求ip地址
	user.Avator = "/static/upload/default/user-default-60x60.png"
	user.Nickname = aul.Nickname
	if err := user.Insert(); err != nil {
		logs.Warning(err.Error())
		this.Data["json"] = RetResource(false, nil, err.Error())
		this.ServeJSON()
		return
	}
	authkey := models.Md5([]byte(this.getClientIp() + "|" + user.Password))
	this.Ctx.SetCookie("auth", strconv.FormatInt(user.Id, 10)+"|"+authkey, 86400)
	//fmt.Println("========FormatInt========", strconv.FormatInt(user.Id, 10))
	this.SetSession("userId", user.Id)
	this.SetSession("userName", user.Username)
	logs.Info(fmt.Sprintf("userid:%d,username:%s,注册成功", user.Id, user.Username))
	this.Data["json"] = RetResource(true, nil, "注册成功")
	this.ServeJSON()
}

func checkUsername(username string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z]([a-zA-Z0-9-_]{4,14})+$", username); !ok {
		return false
	}
	return true
}

func checkPassword(password string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9[:punct:]]{4,19}$", password); !ok {
		return false
	}
	return true
}
