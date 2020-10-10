package models

import (
	"fmt"
	"newSecKill2/common/initall"

	"github.com/astaxie/beego/orm"
	"github.com/beego/admin/src/models"
	"github.com/coreos/etcd/clientv3"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB                  orm.Ormer
	SecKillConf         initall.ConfigAll
	SecKillActivityList []SecKillActivity
	EtcdClient          *clientv3.Client
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

}

func initAll() (err error) {
	if SecKillConf, err = initall.InitConfig(); err != nil {
		return
	}
	if EtcdClient, err = initall.InitEtcd(); err != nil {
		return
	}
	/*
		etcdKey := GetEtcdKey()
		SecKillActivityList, err = loadActivityFromEtcd(etcdKey)
		if err != nil {
			return
		}
	*/
	return
}
