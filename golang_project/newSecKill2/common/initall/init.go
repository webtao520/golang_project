package initall

import (
	//"errors"
	//"fmt"

	//"github.com/astaxie/beego"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"go.etcd.io/etcd/clientv3"
	//"github.com/astaxie/beego/logs"
)

var (
	Db          orm.Ormer
	SecKillConf ConfigAll
	EtcdClient  *clientv3.Client
)

/*
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
*/

func InitEtcd() (EtcdClient *clientv3.Client, err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{SecKillConf.EtcdConfig.Address},
		DialTimeout: time.Duration(SecKillConf.EtcdConfig.DailTimeOut) * time.Second,
	})
	/**
	要将整数个某时间单元表示为Duration类型值
	seconds := 10
	fmt.Print(time.Duration(seconds)*time.Second) // prints 10s
	*/
	if err != nil {
		logs.Error("connect etcd failed, err:", err)
		return
	}
	EtcdClient = cli
	logs.Debug("init etcd success")
	return
}
