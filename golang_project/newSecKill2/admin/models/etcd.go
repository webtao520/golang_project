package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
)

func GetEtcdKey() (etcdKey string) {
	PrefixKey := SecKillConf.EtcdConfig.PrefixKey
	/**
	func HasSuffix(s, suffix string) bool
	判断s是否有后缀字符串suffix。
	*/
	if strings.HasSuffix(PrefixKey, "/") == false {
		PrefixKey = PrefixKey + "/"
	}
	ProductKey := SecKillConf.EtcdConfig.ProductKey
	if strings.HasSuffix(ProductKey, "/") == false {
		ProductKey = ProductKey + "/"
	}
	etcdKey = PrefixKey + ProductKey
	return
}

// 添加，修改，删除 活动到mysql 成功后同步到etcd
func (this *SecKillActivity) syncActivityToEtcd(types string, activity SecKillActivity) (err error) {
	typeMap := map[string]int{"add": 1, "update": 2, "del": 3}
	// 判断是否有效参数
	typeValue, ok := typeMap[types]
	if !ok {
		err = errors.New(fmt.Sprintln("sync Activity To Etcd 参数错误, 错误参数:%v", types))
		logs.Warn(err)
		return
	}
	etcdKey := GetEtcdKey()
	activityList, err := loadActivityFromEtcd(etcdKey)
	switch typeValue {
	case 1:
		activityList = append(activityList, activity)
	default:
		logs.Error("sync Activity To Etcd warning : %v", types)
	}
	jsonActivityList, err := json.Marshal(activityList)

	if err != nil {
		logs.Error("json marshal activityList failed, err : %v", err)
		return
	}
	// 设置超时时间
	putTimeOut := time.Second * time.Duration(SecKillConf.EtcdConfig.PutTimeOut)
	ctx, cancel := context.WithTimeout(context.Background(), putTimeOut)
	defer func(cancel context.CancelFunc) { cancel() }(cancel)
	// 存入ETCD
	stringActivityList := string(jsonActivityList)
	_, err = EtcdClient.Put(ctx, etcdKey, stringActivityList)
	if err != nil {
		logs.Error("put [%s] to etcd [%s] failed, err : %v", stringActivityList, etcdKey, err)
		return
	}
	fmt.Println("etcdKey====>", etcdKey)
	return
}

func loadActivityFromEtcd(etcdKey string) (activityList []SecKillActivity, err error) {
	// 设置超时时间
	getTimeOut := time.Second * time.Duration(SecKillConf.EtcdConfig.GetTimeOut)
	ctx, cancel := context.WithTimeout(context.Background(), getTimeOut)
	defer func(cancel context.CancelFunc) { cancel() }(cancel)
	// 取值
	etcdActivityInfo, err := EtcdClient.Get(ctx, etcdKey)
	if err != nil {
		logs.Error("get [%s] from etcd failed, err : %v", etcdKey, err)
		return
	}
	for _, v := range etcdActivityInfo.Kvs {
		err = json.Unmarshal(v.Value, &activityList)
		if err != nil {
			logs.Error("Unmarshal activityInfo failed, err : %v", err)
			return
		}
	}
	return
}
