package models

import (
	"fmt"
	"secKill/common/initall"
	"sync"

	"github.com/astaxie/beego/orm"
	"github.com/coreos/etcd/clientv3"
	"github.com/garyburd/redigo/redis"
)

var (
	EtcdKey          string
	DB               orm.Ormer
	EtcdClient       *clientv3.Client
	AccessRedisPool  *redis.Pool
	DisposeRedisPool *redis.Pool
	SecKillConf      initall.ConfigAll
	SecKillInfoList  []SecKillInfo
	SecKillInfoMap   = make(map[int]SecKillInfo)
	MutexLock        sync.Mutex
	RWLock           sync.RWMutex
)

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SecKillLuckyUser))

	if err := initAll(); err != nil {
		panic(fmt.Sprintln("init database failed, err:%v", err))
	}

	// 自动建表
	orm.RunSyncdb("default", false, true)

	// 监听 ETCD 变化
	go WatchSecKillEtcd(EtcdKey)
	go ReadAccessRedis()
	go Dispose()
	go WriteDisposeRedis()

}

func initAll() (err error) {
	if SecKillConf, err = initall.InitConfig(); err != nil {
		return
	}
	if DB, err = initall.InitMysql(); err != nil {
		return
	}
	if EtcdClient, err = initall.InitEtcd(); err != nil {
		return
	}
	if AccessRedisPool, err = initall.InitAccessRedis(); err != nil {
		return
	}
	if DisposeRedisPool, err = initall.InitDisposeRedis(); err != nil {
		return
	}

	if err = initall.InitLogs(); err != nil {
		return
	}
	// 初始化	EtcdKey
	EtcdKey = GetEtcdKey()
	// 加载秒杀的信息
	SecKillInfoList, err = GetSecKillInfoListFromEtcd()
	fmt.Println(SecKillInfoList)
	if err != nil {
		return
	}
	return
}
