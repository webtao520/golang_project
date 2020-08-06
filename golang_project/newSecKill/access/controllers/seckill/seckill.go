package seckill

import (
	"fmt"
	"newSecKill/access/models"
	"strings"
	"strconv"
	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

type SecKillController struct {
	beego.Controller
}

var (
	NowSecKillModel  *models.NowSecKillInfo = models.NewNowSecKillModel()
)

func (this *SecKillController) Failed() {
	this.TplName = "seckill/failed.html"
	return
}

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
	secKillInfo:=NowSecKillModel.GetSecKillInfo()
	if  secKillInfo == nil {
		this.Success("/","秒杀活动未开始")
		return
	}
	user := this.GetSession("user")
	//fmt.Println("==========>",user)  // &{10 张涛 29fe7d4261683524a453faea8abe2295 1183681473@qq.com 15021769391 1 1 2020-08-05 10:33:13.9406638 +0800 CST}
	this.Data["secKillInfo"] = secKillInfo
	this.Data["user"] = user
	this.TplName = "seckill/index.html"
	return
}



 //秒杀
func (this *SecKillController) SecKill(){
	ClientAddr:=this.Ctx.Request.Referer()
	//fmt.Println("========>", ClientAddr) // http://localhost:8080/seckill/index
	var ClientRefence string
	if len(ClientAddr) > 0 {
		ClientRefence=strings.TrimLeft(strings.Split(ClientAddr,":")[1],"//") // localhost
	}
	CloseNotify := this.Ctx.ResponseWriter.CloseNotify()
	ActivityId, _ := strconv.Atoi(this.Ctx.Input.Param(":ActivityId"))
	if ActivityId < 1 {
		logs.Error("SecKill get ActivityId err ")
		this.Failed()
		return
	}
	// 获取用户ip
	UserIp:=this.Ctx.Request.RemoteAddr
	// 获取用户UserId
	UserId := this.GetSession("user").(*models.SecKillUser).UserId
	secKillRequest := &models.SecKillRequest{
		UserId:        UserId,  // 5
		Ip:            UserIp, // 18
		ActivityId:    ActivityId, // [::1]:65307
		ClientAddr:    ClientAddr, // http://localhost:8080/seckill/index
		ClientRefence: ClientRefence, // localhost
		CloseNotify:   CloseNotify, // 0xc0004be2a0
	}
	/*
	 ====秒杀请求=== &{5 18 [::1]:65307 http://localhost:8080/seckill/index localhost 0xc0004be2a0}
	 fmt.Println("====秒杀请求===",secKillRequest) 
	*/
	// 处理秒杀结果
	data,err:=models.SecKill(secKillRequest)
	if err != nil {
		this.Failed()
		return
	}
	this.Data["UserId"]=data["UserId"]
	this.Data["token"]=data["token"]
	this.Data["ActivityId"]=data["ActivityId"]
	this.Data["UserName"]=this.GetSession("user").(*models.SecKillUser).UserName
	this.TplName = "seckill/success.html"

}