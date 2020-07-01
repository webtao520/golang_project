package seckill

import (
	"fmt"
	"secKill/access/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

type SecKillController struct {
	beego.Controller
}

var (
	NowSecKillModel *models.NowSecKillInfo = models.NewNowSecKillModel()
	AntispamModel   *models.Antispam       = models.NewAntispamModel()
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

func (this *SecKillController) SecKill() {

	ClientAddr := this.Ctx.Request.Referer()
	var ClientRefence string
	if len(ClientAddr) > 0 {
		ClientRefence = strings.TrimLeft(strings.Split(ClientAddr, ":")[1], "//")
	}

	CloseNotify := this.Ctx.ResponseWriter.CloseNotify()
	ActivityId, _ := strconv.Atoi(this.Ctx.Input.Param(":ActivityId"))
	if ActivityId < 1 {
		logs.Error("SecKill get ActivityId err ")
		this.Failed()
		return
	}
	// 获取用户 ip
	UserIp := this.Ctx.Request.RemoteAddr
	// 获取用户UserId
	UserId := this.GetSession("user").(*models.SecKillUser).UserId

	secKillRequest := &models.SecKillRequest{
		UserId:        UserId,
		Ip:            UserIp,
		ActivityId:    ActivityId,
		ClientAddr:    ClientAddr,
		ClientRefence: ClientRefence,
		CloseNotify:   CloseNotify,
	}
	// 处理秒杀结果
	data, err := models.SecKill(secKillRequest)
	if err != nil {
		this.Failed()
		return
	}
	this.Data["UserId"] = data["UserId"]
	this.Data["token"] = data["token"]
	this.Data["ActivityId"] = data["ActivityId"]
	this.Data["UserName"] = this.GetSession("user").(*models.SecKillUser).UserName

	this.TplName = "seckill/success.html"
}
