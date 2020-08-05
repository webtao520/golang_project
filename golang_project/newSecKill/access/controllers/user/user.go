package user

import (
	"fmt"
	"time"
	"newSecKill/access/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/astaxie/beego/validation"

)

// 定义用户控制器
type UserController struct {
	beego.Controller
}

var (
	cpt       *captcha.Captcha
	valid     =validation.Validation{} // 初始化
	UserModel *models.SecKillUser= models.NewUserModel()
)

// 验证码 初始化
func  init(){
	   // use beego cache system store the captcha data
	   store := cache.NewMemoryCache()
	   cpt = captcha.NewWithFilter("/captcha/", store)
	   cpt.ChallengeNums = 6
	   cpt.StdWidth = 200
	   cpt.StdHeight = 80
}

func (this *UserController) Error(err interface{}){
	// 获取当前页面url 
	url:=this.Ctx.Request.Referer()
	this.Data["error"]=fmt.Sprintln(err)
	this.Redirect(url,302)
	return
}


// 用户注册
func (this *UserController) Register() {
	if this.Ctx.Input.IsPost() {
		UserName := this.GetString("UserName")
		if len(UserName) < 1 {
			this.Error("请输入用户名")
			return
		}
		UserPwd := this.GetString("UserPwd")
		if len(UserPwd) < 6 {
			this.Error("密码不小于6位")
			return
		}
		UserEmail := this.GetString("UserEmail")
		validResult := valid.Email(UserEmail, "UserEmail")
		if !validResult.Ok {
			this.Error(fmt.Sprintln("Email邮箱格式 : ", validResult.Error.Key, validResult.Error.Message))
			return
		}
		UserMobile, err := this.GetUint64("UserMobile")
		if err != nil {
			this.Error(err)
			return
		}
		validResult = valid.Mobile(UserMobile, "UserMobile")
		if !validResult.Ok {
			this.Error(fmt.Sprintln("Mobile手机号 : ", validResult.Error.Key, validResult.Error.Message))
			return
		}
		// 验证输入的验证码
		captcha := cpt.VerifyReq(this.Ctx.Request)
		if !captcha {
			this.Error("验证码有误！")
			return
		}
		user := &models.SecKillUser{
			UserName:   UserName,
			UserPwd:    UserPwd,
			UserEmail:  UserEmail,
			UserMobile: UserMobile,
			Status:     1,
			IsLogin:    1,
			CreateTime: time.Now().Local(),
		}
		_, err = UserModel.InsertUser(user)
		if err != nil {
			this.Error(err)
			return
		}
		// 设置 session
		this.SetSession("user", user)
		// 设置 cookie
		this.Ctx.SetCookie("user_name", UserName)
		this.Ctx.SetCookie("user_pwd", UserPwd)

		this.Redirect("/seckill/index", 302)

	} else {
		this.TplName = "user/register.html"
	}

	return
}