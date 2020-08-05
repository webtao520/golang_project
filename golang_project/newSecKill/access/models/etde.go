package models

import (
	"context"
	"encoding/json"
	//"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
)

/*

loadsecConf
	开始时间，结束时间，状态，库存，剩余库存

	监听etcd变化

	如果etcd有变化，更新etcd
	前后台都在用配置，更改时需要加读写锁

*/

type SecKillInfo struct {
	ActivityId    int
	ActivityName  string
	ProductId     int
	StartTime     time.Time   // 开始时间
	EndTime       time.Time  // 结束时间
	Total         int32 // 活动商品数量
	Status        int // 活动状态（1.未开始，2.进行中，3.已结束）
	SecondLimit   int  // 每秒可售商品个数
	EveryoneLimit int  // 每人限购商品数
	BuyRate       int // 购买到的概率 （百分比）
}


func GetEtcdKey() (EtcdKey string) {
	PrefixKey := SecKillConf.EtcdConfig.PrefixKey
	//fmt.Println("=======>",PrefixKey)
	if strings.HasSuffix(PrefixKey, "/") == false {
		PrefixKey = PrefixKey + "/"
	}
	ProductKey := SecKillConf.EtcdConfig.ProductKey
	if strings.HasSuffix(ProductKey, "/") == false {
		ProductKey = ProductKey + "/"
	}
	EtcdKey = PrefixKey + ProductKey
	return
}

// 加载秒杀信息
func GetSecKillInfoListFromEtcd() (SecKillInfoList []SecKillInfo, err error) {
	// 设置超时时间
	getTimeOut:=time.Second * time.Duration(SecKillConf.EtcdConfig.GetTimeOut) // 20s
	ctx,cancel:=context.WithTimeout(context.Background(),getTimeOut)
	defer func(cancel context.CancelFunc) { cancel() }(cancel)
	// 取值
	etcdSecKillInfo,err:=EtcdClient.Get(ctx,EtcdKey)
    if err !=nil {
		logs.Error("get [%s] from etcd failed, err : %v", EtcdKey, err)
		return
	}
	// func Unmarshal(data []byte, v interface{}) error
	for _,v:=range etcdSecKillInfo.Kvs {
		//fmt.Println("=======>",v.Value)  //[]byte
		//fmt.Printf("%T",v.Value)  //(  []uint8 / byte  )     table `sec_kill_user` already exists, skip
		err=json.Unmarshal(v.Value,&SecKillInfoList)
		if err !=nil {
			logs.Error("Unmarshal etcdSecKillInfo failed, err : %v", err)
			return
		}
		//fmt.Println(SecKillInfoList)
	}
	// []SecKillInfo 转换为map
	SecKillInfoSwitchover(SecKillInfoList)
	return
}

// []SecKillInfo 转换为map
func SecKillInfoSwitchover(SecKillInfoList []SecKillInfo){
	mapInfo:=make(map[int]*SecKillInfo)
	for _,v:=range SecKillInfoList {
		mapInfo[v.ActivityId]=&v
	}
	// SecKillInfoMap     =make(map[int]*SecKillInfo)
	SecKillInfoMap = mapInfo
	//fmt.Println("======etcd 原始数据============>>",SecKillInfoMap)
}