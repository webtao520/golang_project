package models

import (
	"fmt"
	"newSecKill/common/initall"
	"sync"

	"github.com/astaxie/beego/orm"
	"github.com/beego/admin/src/models"
	"github.com/coreos/etcd/clientv3"
)

var (
	DB                  orm.Ormer
	EtcdClient          *clientv3.Client
	SecKillConf         initall.ConfigAll
	SecKillActivityList []SecKillActivity // 活动数据
	ActivityModel       *SecKillActivity  = NewActivityModel()
	MutexLock           sync.Mutex        // 互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个 goroutine 可以访问到共享资源（同一个时刻只有一个线程能够拿到锁）
)

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SecKillProduct), new(SecKillActivity))
	//数据库连接
	models.Connect()
	// 自动建表
	orm.RunSyncdb("default", false, true)

	DB = orm.NewOrm()

	if err := initAll(); err != nil {
		panic(fmt.Sprintln("init database failed, err:%v", err))
	}

	go ActivityModel.WatchActivity()

}

func initAll() (err error) {
	// 初始话配置信息
	if SecKillConf, err = initall.InitConfig(); err != nil {
		return
	}
	// 连接etcd数据库
	if EtcdClient, err = initall.InitEtcd(); err != nil {
		return
	}

	etcdKey := GetEtcdKey()
	SecKillActivityList, err = loadActivityFromEtcd(etcdKey)
	//fmt.Println("============>", SecKillActivityList)
	if err != nil {
		return
	}
	return
}
