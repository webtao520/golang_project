package seckill

import (
	"fmt"
	"newSecKill/access/models"
	"strings"

	"github.com/astaxie/beego"
)

type SecKillController struct {
	beego.Controller
}

var (
	NowSecKillModel *models.NowSecKillInfo = models.NewNowSecKillModel()
)

func (this *SecKillController) Error(err interface{}) {
	fmt.Println(err)
	err = fmt.Sprintln(err)
	this.Data["error"] = err
	url := this.Ctx.Request.Referer()
	this.Redirect(url, 302)
	return
}

func (this *SecKillController) Success(url string, message string) {
	fmt.Println(message)
	this.Data["message"] = message
	this.Redirect(url, 302)
	return
}

func (this *SecKillController) Failed() {
	this.TplName = "seckill/failed.html"
	return
}

func (this *SecKillController) Index() {
	secKillInfo := NowSecKillModel.GetSecKillInfo()
	if secKillInfo == nil {
		this.Success("/", "秒杀活动未开始")
		return
	}
	this.Data["secKillInfo"] = secKillInfo
	user := this.GetSession("user")
	this.Data["user"] = user
	this.TplName = "seckill/index.html"
	return
}

// 用户秒杀
func (this *SecKillController) SecKill() {
	ClientAddr := this.Ctx.Request.Referer() // http://127.0.0.1:8888/seckill/index 返回请求url
	//fmt.Println(ClientAddr)
	var ClientRefence string
	if len(ClientAddr) > 0 {
		ClientRefence = strings.TrimLeft(strings.Split(ClientAddr, ":")[1], "//")
	}
}
