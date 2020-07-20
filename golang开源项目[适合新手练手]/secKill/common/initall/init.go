package initall

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/clientv3"
	"github.com/garyburd/redigo/redis"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db               orm.Ormer
	EtcdClient       *clientv3.Client
	SecKillConf      ConfigAll
	AccessRedisPool  *redis.Pool
	DisposeRedisPool *redis.Pool
)

func InitAll() (SecKillConf ConfigAll, err error) {

	SecKillConf, err = InitConfig()
	if err != nil {
		err = errors.New(fmt.Sprintf("init config err : ", err))
		return
	}

	Db, err = InitMysql()
	if err != nil {
		err = errors.New(fmt.Sprintf("init mysql err : ", err))
		return
	}

	EtcdClient, err = InitEtcd()
	if err != nil {
		err = errors.New(fmt.Sprintf("init etcd err : ", err))
		return
	}
	AccessRedisPool, err = InitAccessRedis()
	if err != nil {
		return
	}
	DisposeRedisPool, err = InitDisposeRedis()
	if err != nil {
		return
	}
	err = InitLogs()
	if err != nil {
		err = errors.New(fmt.Sprintf("init logs err : ", err))
		return
	}
	return
}

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

func InitLogs() (err error) {
	logPath := SecKillConf.LogConfig.LogPath
	config := fmt.Sprintf(`{"filename":"%s"}`, logPath)
	// 日志配置
	beego.SetLogger("file", config)
	// 设置级别
	beego.SetLevel(beego.LevelDebug)
	// 输出文件名和行号
	beego.SetLogFuncCall(true)

	logs.Debug("init logs success")
	return
}

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

func InitDisposeRedis() (pool *redis.Pool, err error) {
	pool = &redis.Pool{ //实例化一个连接池
		MaxIdle:     SecKillConf.DisposeRedisConfig.MaxIdle,
		MaxActive:   SecKillConf.DisposeRedisConfig.MaxActive,
		IdleTimeout: time.Duration(SecKillConf.DisposeRedisConfig.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", SecKillConf.DisposeRedisConfig.Address)
		},
	}
	conn := pool.Get()
	defer func(conn redis.Conn) { conn.Close() }(conn)
	_, err = conn.Do("ping")
	if err != nil {
		err = errors.New(fmt.Sprintf("ping DisposeRedis failed, err : %v", err))
		logs.Error(err)
		return
	}
	return
}

func InitBlacklistRedis() (pool *redis.Pool, err error) {
	pool = &redis.Pool{ //实例化一个连接池
		MaxIdle:     SecKillConf.BlacklistRedisConfig.MaxIdle,
		MaxActive:   SecKillConf.BlacklistRedisConfig.MaxActive,
		IdleTimeout: time.Duration(SecKillConf.BlacklistRedisConfig.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", SecKillConf.BlacklistRedisConfig.Address)
		},
	}
	conn := pool.Get()
	defer func(conn redis.Conn) { conn.Close() }(conn)
	_, err = conn.Do("ping")
	if err != nil {
		err = errors.New(fmt.Sprintf("ping BlacklistRedisConfig failed, err : %v", err))
		logs.Error(err)
		return
	}
	return
}
