package activity

import (
	"errors"
	"fmt"
	"secKill/admin/models"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

var ActivityModel *models.SecKillActivity = models.NewActivityModel()

type ActivityController struct {
	beego.Controller
}

func (this *ActivityController) Error(err interface{}) {
	// 获取当前页面url
	url := this.Ctx.Request.Referer()
	this.Data["error"] = fmt.Sprintln(err)
	fmt.Println(err)
	this.Redirect(url, 302)
	return
}

func (this *ActivityController) Success(url string, msg interface{}) {
	fmt.Println(msg)
	this.Data["success"] = fmt.Sprintln(msg)
	this.Redirect(url, 302)
	return
}

func (this *ActivityController) Index() {
	ActivityList, err := ActivityModel.GetActivityList(0, 10)
	if err != nil {
		this.Error(err)
		return
	}
	this.Data["ActivityList"] = ActivityList
	this.TplName = "seckill/activity/index.html"
	return
}

// 获取HTML提交的数据 add update 通用
func (this *ActivityController) getDataFormHtml() (activity models.SecKillActivity, err error) {
	ActivityName := this.GetString("ActivityName")
	if len(ActivityName) <= 0 {
		err = errors.New("请输入活动名称")
		return
	}
	ProductId, err := this.GetInt("ProductId")
	if err != nil {
		err = errors.New("请选择活动商品")
		return
	}
	receiveStartTime := this.GetString("StartTime")
	if len(receiveStartTime) <= 0 {
		err = errors.New("请选择开始时间")
		return
	}
	StartTime, err := time.ParseInLocation("2006-01-02 15:04:05", receiveStartTime, time.Local)
	if err != nil {
		logs.Warn("switch time receiveStartTime err : %v", err)
		return
	}
	fmt.Println(StartTime)
	receiveEndTime := this.GetString("EndTime")
	if len(receiveEndTime) <= 0 {
		err = errors.New("请选择结束时间")
		return
	}
	EndTime, err := time.ParseInLocation("2006-01-02 15:04:05", receiveEndTime, time.Local)
	if err != nil {
		logs.Warn("switch time receiveEndTime err : %v", err)
		return
	}
	// time.Time 比较
	if !StartTime.Before(EndTime) {
		err = errors.New("活动开始时间必须早于结束时间")
		return
	}
	if StartTime.Before(time.Now().Local()) {
		err = errors.New("活动开始时间必须晚于当前时间")
		return
	}
	Total, err := this.GetInt("Total")
	if err != nil {
		err = errors.New("Total 必须是有效数字")
		return
	}
	Status, err := this.GetInt("Status")
	if err != nil {
		err = errors.New("Status 必须是有效数字")
		return
	}
	SecondLimit, err := this.GetInt("SecondLimit")
	if err != nil {
		err = errors.New("SecondLimit 必须是有效数字")
		return
	}
	EveryoneLimit, err := this.GetInt("EveryoneLimit")
	if err != nil {
		err = errors.New("EveryoneLimit 必须是有效数字")
		return
	}
	BuyRate, err := this.GetInt("BuyRate")
	if err != nil {
		return
	}
	activity = models.SecKillActivity{
		ActivityName:  ActivityName,
		ProductId:     ProductId,
		StartTime:     StartTime,
		EndTime:       EndTime,
		Total:         Total,
		Status:        Status,
		SecondLimit:   SecondLimit,
		EveryoneLimit: EveryoneLimit,
		BuyRate:       BuyRate,
	}
	return
}

func (this *ActivityController) AddActivity() {
	if this.Ctx.Input.IsPost() {
		activity, err := this.getDataFormHtml()
		if err != nil {
			this.Error(err)
			return
		}
		_, err = ActivityModel.InsertActivity(&activity)
		if err != nil {
			this.Error(err)
			return
		}
		this.Success("index", "添加成功！")
	} else {
		ProductList, err := models.NewProductModel().GetProductList(0, 10)
		if err != nil {
			this.Error(err)
			return
		}
		this.Data["product"] = ProductList
		this.TplName = "seckill/activity/add.html"
	}
	return
}

func (this *ActivityController) UpdateActivity() {
	ActivityId, err := this.GetInt("ActivityId")
	if err != nil {
		err = errors.New(fmt.Sprintln("获取 ActivityId err : %v ", err))
		this.Error(err)
		return
	}
	if this.Ctx.Input.IsPost() {
		activity, err := this.getDataFormHtml()
		if err != nil {
			this.Error(err)
			return
		}
		activity.ActivityId = ActivityId
		_, err = ActivityModel.UpdateActivity(&activity)
		if err != nil {
			this.Error(err)
			return
		}
		this.Success("index", "修改成功！")
		return
	} else {
		activity, err := ActivityModel.GetActivityById(ActivityId)
		if err != nil {
			this.Error(err)
			return
		}

		ProductList, err := models.NewProductModel().GetProductList(0, 10)
		if err != nil {
			this.Error(err)
			return
		}
		type ActivityInfo struct {
			Activity    models.SecKillActivity
			ProductList []models.SecKillProduct
		}
		activityInfo := ActivityInfo{
			Activity:    activity,
			ProductList: ProductList,
		}
		this.Data["activityInfo"] = activityInfo
		this.TplName = "seckill/activity/edit.html"
	}
	return
}

func (this *ActivityController) DelActivity() {
	ActivityId, err := this.GetInt("ActivityId")
	if err != nil {
		this.Error(err)
		return
	}
	activity, err := ActivityModel.GetActivityById(ActivityId)
	if err != nil {
		this.Error(err)
		return
	}
	// 进行中 和 已结束 的活动不能删除
	if activity.Status != 1 {
		this.Error("进行中 和 已结束 的活动不能删除")
		return
	}
	_, err = ActivityModel.DelActivity(&activity)
	if err != nil {
		this.Error(err)
		return
	}
	this.Success("index", "删除成功！")
	return
}
