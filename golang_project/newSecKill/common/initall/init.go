package initall

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/coreos/etcd/clientv3"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db          orm.Ormer
	EtcdClient  *clientv3.Client
	SecKillConf ConfigAll
)

/*
func InitAll() (SecKillConf ConfigAll, err error) {
	SecKillConf, err = InitConfig() // 数据库基础配置结构赋值
	if err != nil {
		err = errors.New(fmt.Sprintf("init config err : ", err))
		return
	}

	Db, err = InitMysql()
	if err != nil {
		err = errors.New(fmt.Sprintf("init mysql err : ", err))
		return
	}
	return
}
*/

func InitMysql() (Db orm.Ormer, err error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", SecKillConf.MysqlConfig.UserName, SecKillConf.MysqlConfig.PassWd,
		SecKillConf.MysqlConfig.Host, SecKillConf.MysqlConfig.Port, SecKillConf.MysqlConfig.DbName)
	// 链接数据库
	err = orm.RegisterDataBase("default", "mysql", dns)

	if err != nil {
		errors.New(fmt.Sprintf("connect mysql err : ", err))
		return
	}

	// 创建一个数据库链接 Ormer
	Db = orm.NewOrm()

	logs.Debug("connect mysql success")

	return
}

/**
  创建client
*/
func InitEtcd() (EtcdClient *clientv3.Client, err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{SecKillConf.EtcdConfig.Address},
		DialTimeout: time.Duration(SecKillConf.EtcdConfig.DailTimeOut) * time.Second,
	})

	if err != nil {
		logs.Error("connect etcd failed, err:", err)
		return
	}
	EtcdClient = cli
	logs.Debug("init etcd success")
	return
}

// redis 配置
func InitAccessRedis() (pool *redis.Pool, err error) {
	pool = &redis.Pool{ //实例化一个连接池
		MaxIdle:     SecKillConf.AccessRedisConfig.MaxIdle,
		MaxActive:   SecKillConf.AccessRedisConfig.MaxActive,
		IdleTimeout: time.Duration(SecKillConf.AccessRedisConfig.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", SecKillConf.AccessRedisConfig.Address)
		},
	}
	conn := pool.Get()
	defer func(conn redis.Conn) { conn.Close() }(conn)
	_, err = conn.Do("ping")
	if err != nil {
		err = errors.New(fmt.Sprintf("ping AccessRedis failed, err : %v", err))
		logs.Error(err)
		return
	}
	return
}
