package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
)

// 根据 req 判断 ,是否恶意用户 ，是否超出活动时间
func AntispamByActivityId(req SecKillRequest) (err error) {
	// 用户访问ip判断
	if req.ClientRefence != "127.0.0.1" {
		err = fmt.Errorf("非法ip : %v 访问 ", req.ClientRefence)
		logs.Error(err)
		return
	}

	Activity, ok := SecKillInfoMap[req.ActivityId]
	if !ok {
		err = fmt.Errorf("ActivityId : %d 没找到 Activity 信息", req.ActivityId)
		fmt.Println(err)
		return
	}
	if Activity.Status != 2 {
		err = fmt.Errorf("ActivityId : %d ,Activity.Status != 2", req.ActivityId)
		fmt.Println(err)
		return
	}
	now := time.Now().Local()
	if !now.After(Activity.StartTime) {
		err = fmt.Errorf("Activity.StartTime 异常 , StartTime : %v , NOW : %v", Activity.StartTime, now)
		fmt.Println(err)
		return
	}
	if !Activity.EndTime.After(now) {
		err = fmt.Errorf("Activity.EndTime 异常 , EndTime : %v , NOW : %v", Activity.EndTime, now)
		fmt.Println(err)
		return
	}
	return
}
