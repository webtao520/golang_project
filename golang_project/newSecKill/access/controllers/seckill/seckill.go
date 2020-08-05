package seckill

import (
	//"fmt"
	"fmt"
	"newSecKill/access/models"

	"github.com/astaxie/beego"
)

type SecKillController struct {
	beego.Controller
}

var (
	NowSecKillModel  *models.NowSecKillInfo = models.NewNowSecKillModel()
)

func (this *SecKillController) Success(url string, message string){
	//fmt.Println(message)
	this.Data["message"]=message
	this.Redirect(url,302)
	return
}

func (this *SecKillController) Error(err interface{}) {
	//fmt.Println(err)
	err = fmt.Sprintln(err)
	this.Data["error"] = err
	url := this.Ctx.Request.Referer()
	this.Redirect(url, 302)
	return
}

func (this *SecKillController) Index() {
	seckillInfo:=NowSecKillModel.GetSecKillInfo()
	if  seckillInfo == nil {
		this.Success("/","秒杀活动未开始")
		return
	}
	user := this.GetSession("user")
	this.Data["user"] = user
	this.TplName = "seckill/index.html"
	return
}