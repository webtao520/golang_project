package seckill

import (
	"fmt"
	"newSecKill/access/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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
		ClientRefence = strings.TrimLeft(strings.Split(ClientAddr, ":")[1], "//") //获取主机地址
	}
	CloseNotify := this.Ctx.ResponseWriter.CloseNotify()               // golang捕获http.ResponseWriter被close
	ActivityId, _ := strconv.Atoi(this.Ctx.Input.Param(":ActivityId")) //字符串转换为int 按给定的键返回路由器参数。
	if ActivityId < 1 {
		logs.Error("SecKill get ActivityId err")
		this.Failed()
		return
	}
	// 获取客户端的地址和端口号
	UserIp := this.Ctx.Request.RemoteAddr // 127.0.0.1:55512
	// 获取用户userid
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
