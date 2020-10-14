package models

import "time"

//"fmt"

//"fmt"

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
	CloseNotify   <-chan bool         `json:"_"`
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
	//data = make(map[string]interface{})
	//fmt.Println(req) // &{3 12 127.0.0.1:57261 http://127.0.0.1:8888/seckill/index 127.0.0.1 0xc0005a40e0 <nil>}
	t := time.NewTicker(time.Second * 10)
	defer func(t *time.Ticker) {
		t.Stop()
	}(t)

	return
}

// 获取秒杀信息
func (this *NowSecKillInfo) GetSecKillInfo() (nowSecKillInfo []NowSecKillInfo) {
	secKillInfo := NowSecKillInfo{}
	// fmt.Println("===============>",SecKillInfoMap)    ===============> map[3:0xc0001a0300]
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
