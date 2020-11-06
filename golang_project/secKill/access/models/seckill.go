package models

import (
	"fmt"
	"sync/atomic"
	"time"
)

// 当前的秒杀信息
type NowSecKillInfo struct {
	ActivityId   int    //活动id
	ActivityName string //活动名称
	Total        int32  //活动商品数量
}

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

func NewNowSecKillModel() *NowSecKillInfo {
	return &NowSecKillInfo{}
}

func (this *NowSecKillInfo) GetSecKillInfo() (nowSecKillInfo []NowSecKillInfo) {
	// now := time.Now().Local()
	secKillInfo := NowSecKillInfo{}
	fmt.Println(SecKillInfoMap)
	for _, v := range SecKillInfoMap {
		if v.Status != 2 || v.Total < 1 {
			continue
		}
		secKillInfo.ActivityId = v.ActivityId
		secKillInfo.ActivityName = v.ActivityName
		secKillInfo.Total = v.Total
		nowSecKillInfo = append(nowSecKillInfo, secKillInfo)
		fmt.Println(nowSecKillInfo)
	}
	return
}

var (
	SecKillRequestChan chan *SecKillRequest           = make(chan *SecKillRequest, 10000)
	SecKillUserMap     map[string]chan *SecKillResult = make(map[string]chan *SecKillResult, 10000)
)

type SecKillRequest struct {
	ActivityId    int
	UserId        int
	Ip            string
	ClientAddr    string
	ClientRefence string
	CloseNotify   <-chan bool `json:"-"`

	ResultChan chan *SecKillResult `json:"-"`
}

type SecKillResult struct {
	ActivityId int
	UserId     int
	Token      string
	Error      error
}

func SecKill(req *SecKillRequest) (data map[string]interface{}, err error) {
	data = make(map[string]interface{})

	err = NewAntispamModel().SecKillAntispam(req.UserId, req.ActivityId, req.Ip)
	if err != nil {
		return
	}
	t := time.NewTicker(time.Second * 10)
	defer func(t *time.Ticker) {
		t.Stop()
	}(t)

	UserKey := fmt.Sprintln(req.ActivityId, "seckill", req.UserId)

	req.ResultChan = make(chan *SecKillResult, 1)
	SecKillUserMap[UserKey] = req.ResultChan
	SecKillRequestChan <- req

	// 总数减 1
	atomic.AddInt32(&SecKillInfoMap[req.ActivityId].Total, -1)

	select {
	case <-t.C:
		fmt.Println("req.ResultChan : ", req.ResultChan)
		err = fmt.Errorf("超时")
		return
	case <-req.CloseNotify:
		err = fmt.Errorf("client already closed")
		return
	case result := <-req.ResultChan:
		err = result.Error
		if err != nil {
			return
		}

		data["ActivityId"] = result.ActivityId
		data["token"] = result.Token
		data["UserId"] = result.UserId
		return
	}
}
