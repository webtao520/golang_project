package models

import (
	"fmt"
	"newSecKill/common/initall"
	"sync"

	"github.com/astaxie/beego/orm"
	"github.com/coreos/etcd/clientv3"
	"github.com/garyburd/redigo/redis"
)

var (
	EtcdKey         string
	DB              orm.Ormer
	SecKillConf     initall.ConfigAll
	EtcdClient      *clientv3.Client
	SecKillInfoMap  = make(map[int]*SecKillInfo)
	SecKillInfoList []SecKillInfo //etcd 配置修改
	//AccessRedisPool    *redis.Pool
	BlacklistRedisPool *redis.Pool
	MutexLock          sync.Mutex
	RWLock             sync.RWMutex
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

func initAll() (err error) {
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

	//加载redis
	// if AccessRedisPool, err = initall.InitAccessRedis(); err != nil {
	// 	return
	// }

	if BlacklistRedisPool, err = initall.InitBlacklistRedis(); err != nil {
		return
	}

	// 初始化 EtcdKey
	EtcdKey = GetEtcdKey()
	//加载秒杀信息
	SecKillInfoList, err = GetSecKillInfoListFromEtcd()
	//fmt.Println("秒杀信息======>",SecKillInfoList)
	if err != nil {
		return
	}
	return
}
