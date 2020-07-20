package models

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/mvcc/mvccpb"
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
	StartTime     time.Time
	EndTime       time.Time
	Total         int32
	Status        int
	SecondLimit   int
	EveryoneLimit int
	BuyRate       int
}

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

// 监听 ETCD 变化
func WatchSecKillEtcd(EtcdKey string) {
	var err error
	for {
		// watch key 监听节点
		var (
			secKillInfoList []SecKillInfo
			watchSuccess    = true
		)
		WatchChan := EtcdClient.Watch(context.Background(), EtcdKey)
		for WatchResponse := range WatchChan {
			for _, Event := range WatchResponse.Events {
				if Event.Type == mvccpb.DELETE {
					logs.Warn("key[%s] 's config deleted", EtcdKey)
					continue
				}
				if Event.Type == mvccpb.PUT && string(Event.Kv.Key) == EtcdKey {
					err = json.Unmarshal(Event.Kv.Value, &secKillInfoList)
					if err != nil {
						logs.Error("key [%s], Unmarshal[%s], err:%v ", err)
						watchSuccess = false
						continue
					}
				}
			}
			if watchSuccess {
				logs.Debug("get config from etcd succ, %v", secKillInfoList)
				MutexLock.Lock()
				SecKillInfoList = secKillInfoList
				// 将 []SecKillInfo 转成 map
				SecKillInfoSwitchover(secKillInfoList)
				MutexLock.Unlock()
			}
		}
	}
}

// 将 []SecKillInfo 转成 map
func SecKillInfoSwitchover(secKillInfoList []SecKillInfo) {
	mapInfo := make(map[int]*SecKillInfo)
	for _, v := range secKillInfoList {
		mapInfo[v.ActivityId] = &v
	}
	SecKillInfoMap = mapInfo
}
