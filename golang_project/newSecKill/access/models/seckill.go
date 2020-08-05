package models

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

func (this *NowSecKillInfo) GetSecKillInfo()(nowSecKillInfo []NowSecKillInfo){
	
}
