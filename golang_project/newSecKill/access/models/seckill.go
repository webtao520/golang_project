package models

import (
	//"fmt"
)

/*
	ActivityId    int    //活动id
	ActivityName  string //活动名称
	ProductId     int    //商品id
	StartTime     int    //活动开始时间
	EndTime       int    //活动结束时间
	Total         int    //活动商品数量
	Status        int    //活动状态（1.未开始，2.进行中，3.已结束）
	SecondLimit   int    //每秒可售商品个数
	EveryoneLimit int    //每人限购商品数
	BuyRate       int    //购买到的概率 （百分比）
*/

// 当前的秒杀信息
type NowSecKillInfo struct {
	ActivityId   int    // 活动id
	ActivityName string // 活动名称
	Total        int32  // 活动商品数量
}

// 初始化秒杀信息models
func NewNowSecKillModel() *NowSecKillInfo {
	return &NowSecKillInfo{}
}


func (this *NowSecKillInfo) GetSecKillInfo() (nowSecKillInfo []NowSecKillInfo) {
	// now := time.Now().Local()
	secKillInfo := NowSecKillInfo{}
	//fmt.Println("原始切片默认值", nowSecKillInfo)
    //fmt.Println(SecKillInfoMap)
	for _, v := range SecKillInfoMap {
		if v.Status != 2 || v.Total < 1 {
			continue
		}
		secKillInfo.ActivityId = v.ActivityId
		secKillInfo.ActivityName = v.ActivityName
		secKillInfo.Total = v.Total
		nowSecKillInfo = append(nowSecKillInfo, secKillInfo)
		//fmt.Println("==========nowSecKillInfo=============>",nowSecKillInfo)
	}
	return
}

type  SecKillRequest struct {
	ActivityId int 
	UserId int 
	Ip string
	ClientAddr    string
	ClientRefence string
	CloseNotify   <-chan bool `json:"-"`
}
