package models

import (
	"fmt"
	"newSecKill/common/initall" 
	"github.com/coreos/etcd/clientv3"
	"github.com/astaxie/beego/orm"
)

var (
	DB orm.Ormer
	SecKillConf  initall.ConfigAll
	EtcdClient         *clientv3.Client
)

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SecKillUser))

	if err := initAll(); err != nil {
		panic(fmt.Sprintln("init database failed, err:%v", err))
	}

	// 自动建表
	orm.RunSyncdb("default", false, true)

	// 监听 ETCD 变化
	//go WatchSecKillEtcd(EtcdKey)

}

func initAll()(err error){
	if SecKillConf, err = initall.InitConfig(); err != nil {
		return
	}
	//fmt.Println("================>",SecKillConf)
	if DB, err = initall.InitMysql(); err != nil {
		return
	}
	if EtcdClient, err = initall.InitEtcd(); err != nil {
		return
	}

	return

}