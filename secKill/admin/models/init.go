package models

import (
	"fmt"
	"secKill/common/initall"
	"sync"

	"github.com/astaxie/beego/orm"
	"github.com/beego/admin/src/models"
	"github.com/coreos/etcd/clientv3"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB                  orm.Ormer
	EtcdClient          *clientv3.Client
	SecKillConf         initall.ConfigAll
	SecKillActivityList []SecKillActivity
	MutexLock           sync.Mutex
	ActivityModel       *SecKillActivity = NewActivityModel()
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
	if SecKillConf, err = initall.InitConfig(); err != nil {
		return
	}

	if EtcdClient, err = initall.InitEtcd(); err != nil {
		return
	}
	if err = initall.InitLogs(); err != nil {
		return
	}
	etcdKey := GetEtcdKey()
	SecKillActivityList, err = loadActivityFromEtcd(etcdKey)
	if err != nil {
		return
	}
	return
}
