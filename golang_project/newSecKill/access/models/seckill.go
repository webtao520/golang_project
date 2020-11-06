package models

import (
	"fmt"
	"sync/atomic"
	"time"
)

var (
	// 通道在声明后，要分配内存
	SecKillRequestChan chan *SecKillRequest           = make(chan *SecKillRequest, 10000)
	SecKillUserMap     map[string]chan *SecKillResult = make(map[string]chan *SecKillResult, 10000)
)

// 当前的秒杀信息
type NowSecKillInfo struct {
	ActivityId   int    // 活动id
	ActivityName string // 活动名称
	Total        int32  // 活动商品数量
}

// 实例化秒杀信息model
func NewNowSecKillModel() *NowSecKillInfo {
	return &NowSecKillInfo{}
}

type SecKillRequest struct {
	ActivityId    int
	UserId        int
	Ip            string
	ClientAddr    string
	ClientRefence string
	CloseNotify   <-chan bool         `json:"_"` // 单项通道，只能读取通道
	ResultChan    chan *SecKillResult `json:"_"`
}

type SecKillResult struct {
	ActivityId int
	UserId     int
	Token      string
	Error      error
}

// 秒杀
func SecKill(req *SecKillRequest) (data map[string]interface{}, err error) {
	data = make(map[string]interface{})

	/*
		err = NewAntispamModel().SecKillAntispam(req.UserId, req.ActivityId, req.Ip)
		if err != nil {
			return
		}
	*/

	t := time.NewTicker(time.Second * 10)
	defer func(t *time.Ticker) {
		t.Stop()
	}(t)

	UserKey := fmt.Sprintln(req.ActivityId, "seckill", req.UserId)

	req.ResultChan = make(chan *SecKillResult, 1)
	SecKillUserMap[UserKey] = req.ResultChan
	SecKillRequestChan <- req

	// 总数减 1  理解下面 逻辑含义
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

// 获取秒杀信息
func (this *NowSecKillInfo) GetSecKillInfo() (nowSecKillInfo []NowSecKillInfo) {
	secKillInfo := NowSecKillInfo{}
	fmt.Println("===============>", SecKillInfoMap) // map[9:0xc00041a180 10:0xc00041a180 11:0xc00041a180]
	for _, v := range SecKillInfoMap {
		if v.Status != 2 || v.Total < 1 {
			continue
		}

		secKillInfo.ActivityId = v.ActivityId
		secKillInfo.ActivityName = v.ActivityName
		secKillInfo.Total = v.Total
		nowSecKillInfo = append(nowSecKillInfo, secKillInfo)
	}
	return
}
