package models

import (
	"context"
	"encoding/json"
	"strings"
	"time"
	//"fmt"

	"github.com/astaxie/beego/logs"
)

/*
CREATE TABLE `sec_kill_activity` (
  `activity_id` int(11) NOT NULL AUTO_INCREMENT,
  `activity_name` varchar(60) NOT NULL DEFAULT '',
  `product_id` int(11) NOT NULL DEFAULT '0',
  `start_time` datetime NOT NULL,
  `end_time` datetime NOT NULL,
  `total` int(11) NOT NULL DEFAULT '0',
  `status` int(11) NOT NULL DEFAULT '2',
  `second_limit` int(11) NOT NULL DEFAULT '0',
  `everyone_limit` int(11) NOT NULL DEFAULT '0',
  `buy_rate` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`activity_id`),
  UNIQUE KEY `activity_name` (`activity_name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

 [{3 秒杀了 3 2020-10-13 16:14:09 +0800 CST 2020-10-17 16:14:09 +0800 CST 500 2 100 50 90}]
*/

// 秒杀活动结构体
type SecKillInfo struct {
	ActivityId    int
	ActivityName  string
	ProductId     int
	StartTime     time.Time
	EndTime       time.Time
	Total         int32
	Status        int
	SecondLimit   int
	EveryoneLimit int
	BuyRate       int
}

// 获取etcd key
func GetEtcdKey() (EtcdKey string) {
	PrefixKey := SecKillConf.EtcdConfig.PrefixKey
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

// 加载秒杀的信息
func GetSecKillInfoListFromEtcd() (SecKillInfoList []SecKillInfo, err error) {

	// 设置超时时间
	getTimeOut := time.Second * time.Duration(SecKillConf.EtcdConfig.GetTimeOut)
	ctx, cancel := context.WithTimeout(context.Background(), getTimeOut)
	defer func(cancel context.CancelFunc) { cancel() }(cancel)
	// 取值
	etcdSecKillInfo, err := EtcdClient.Get(ctx, EtcdKey)
	if err != nil {
		logs.Error("get [%s] from etcd failed, err : %v", EtcdKey, err)
		return
	}
	for _, v := range etcdSecKillInfo.Kvs {
		err = json.Unmarshal(v.Value, &SecKillInfoList)
		if err != nil {
			logs.Error("Unmarshal etcdSecKillInfo failed, err : %v", err)
			return
		}
	}
	// 将 []SecKillInfo 转成 map
	SecKillInfoSwitchover(SecKillInfoList)
	return
}


// 将 []SecKillInfo 转成 map
func SecKillInfoSwitchover(secKillInfoList []SecKillInfo) {
	mapInfo := make(map[int]*SecKillInfo)
	for _, v := range secKillInfoList {
		mapInfo[v.ActivityId] = &v
	}
	//  map[3:0xc00005a000 4:0xc00005a000]
	SecKillInfoMap = mapInfo
}
