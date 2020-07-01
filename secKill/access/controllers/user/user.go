package user

import (
	"fmt"
	"secKill/access/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	beego.Controller
}

var (
	cpt       *captcha.Captcha
	valid                         = validation.Validation{}
	UserModel *models.SecKillUser = models.NewUserModel()
)

func init() {
	// use beego cache system store the captcha data
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 6
	cpt.StdWidth = 200
	cpt.StdHeight = 80
}

func (this *UserController) Error(err interface{}) {
	fmt.Println(err)
	// 获取当前页面url
	url := this.Ctx.Request.Referer()
	this.Data["error"] = fmt.Sprintln(err)
	this.Redirect(url, 302)
	return
}

// 登录
func (this *UserController) Login() {
	if this.Ctx.Input.IsGet() {
		// 获取 session
		userInfo := this.GetSession("user")

		if userInfo != nil {
			// 重定向
			this.Redirect("index", 302)
		} else {
			// 获取 cookie
			this.Data["UserName"] = this.Ctx.GetCookie("UserName")
			this.Data["UserPwd"] = this.Ctx.GetCookie("UserPwd")

			this.TplName = "user/login.html"
		}
	} else {
		UserName := this.GetString("UserName")
		if len(UserName) < 1 {
			this.Error("请输入用户名")
			return
		}
		UserPwd := this.GetString("UserPwd")
		if len(UserPwd) < 6 {
			this.Error("密码长度不小于6位")
			return
		}
		// 验证输入的验证码
		captcha := cpt.VerifyReq(this.Ctx.Request)
		if !captcha {
			this.Error("验证码有误！")
			return
		}
		user, err := UserModel.GetUserByNameAndPwd(UserName, UserPwd)
		if err != nil {
			this.Error(err)
			return
		}
		// 设置 session
		this.SetSession("user", user)
		// 设置 cookie
		this.Ctx.SetCookie("UserName", UserName)
		this.Ctx.SetCookie("UserPwd", UserPwd)

		this.Redirect("/seckill/index", 302)

	}
	return
}

// 注册
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

func (this *UserController) Exit() {
	// 清空 session ，清空后 key 对应的 session value 是 nil
	this.DelSession("user")
	// this.Data["json"] = nil
	// this.ServeJSON()
	this.Redirect("/seckill/user/login", 302)
}
