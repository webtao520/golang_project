package initall

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

// 加载mysql数据库配置信息
type MysqlConfig struct {
	UserName string
	PassWd   string
	Port     int
	DbName   string
	Host     string
}

type EtcdConfig struct {
	Address     string
	PrefixKey   string
	ProductKey  string
	DailTimeOut int
	PutTimeOut  int
	GetTimeOut  int
}

// 系统所有配置信息
type ConfigAll struct {
	MysqlConfig
	EtcdConfig
}

var conf config.Configer

func InitConfig() (ConfigAll ConfigAll, err error) {
	conf, err = config.NewConfig("ini", "../common/conf/common.conf")
	if err != nil {
		err = errors.New(fmt.Sprintf("new config failed, err:", err))
		return
	}

	// mysql 数据库配置信息
	MysqlConfig, err := GetMysqlConfig()
	if err != nil {
		return
	}
	ConfigAll.MysqlConfig = MysqlConfig

	//etcd 配置信息
	EtcdConfig, err := GetEtcdConfig()
	if err != nil {
		return
	}
	ConfigAll.EtcdConfig = EtcdConfig
	SecKillConf = ConfigAll
	logs.Debug("init config success")
	return
}

func GetEtcdConfig() (EtcdConfig EtcdConfig, err error) {
	Address := conf.String("etcd::etcd_addr")
	if len(Address) == 0 {
		err = errors.New("load config of etcd_addr failed , is null")
		logs.Error(err)
		return
	}
	EtcdConfig.Address = Address

	PrefixKey := conf.String("etcd::etcd_prefix_key")
	if len(PrefixKey) == 0 {
		logs.Error("load config of etcd_prefix_key , is null")
		err = errors.New("load config of etcd_prefix_key , is null")
		return
	}
	EtcdConfig.PrefixKey = PrefixKey

	ProductKey := conf.String("etcd::etcd_product_key")
	if len(ProductKey) == 0 {
		logs.Error("load config of etcd_product_key , is null")
		err = errors.New("load config of etcd_product_key , is null")
		return
	}
	EtcdConfig.ProductKey = ProductKey

	DailTimeOut, err := conf.Int("etcd::etcd_dail_timeout")
	if err != nil {
		logs.Error("load config of etcd_dail_timeout err : ", err)
		err = errors.New(fmt.Sprintf("load config of etcd_dail_timeout err : ", err))
		return
	}
	EtcdConfig.DailTimeOut = DailTimeOut

	PutTimeOut, err := conf.Int("etcd::etcd_put_timeout")
	if err != nil {
		logs.Error("load config of etcd_put_timeout err : ", err)
		err = errors.New(fmt.Sprintf("load config of etcd_put_timeout err : ", err))
		return
	}
	EtcdConfig.PutTimeOut = PutTimeOut

	GetTimeOut, err := conf.Int("etcd::etcd_get_timeout")
	if err != nil {
		logs.Error("load config of etcd_get_timeout err : ", err)
		err = errors.New(fmt.Sprintf("load config of etcd_get_timeout err : ", err))
		return
	}
	EtcdConfig.GetTimeOut = GetTimeOut
	return
}

func GetMysqlConfig() (MysqlConfig MysqlConfig, err error) {
	UserName := conf.String("mysql::mysql_user_name")
	if len(UserName) == 0 {
		logs.Error("load config of mysql_user_name failed, is null")
		err = errors.New("load config of mysql_user_name failed, is null")
		return
	}
	MysqlConfig.UserName = UserName

	PassWd := conf.String("mysql::mysql_pass")
	if len(PassWd) == 0 {
		logs.Error("load config of mysql_pass failed, is null")
		err = errors.New("load config of mysql_pass failed, is null")
		return
	}

	MysqlConfig.PassWd = PassWd

	Port, err := conf.Int("mysql::mysql_port")
	if err != nil {
		logs.Error("load config of mysql_port failed, is err:", err)
		err = errors.New(fmt.Sprintf("load config of mysql_port failed, is err: ", err))
		return
	}

	MysqlConfig.Port = Port

	Host := conf.String("mysql::mysql_host")
	if len(Host) == 0 {
		logs.Error("load config of mysql_host failed , is null")
		err = errors.New("load config of mysql_host failed , is null")
		return
	}
	MysqlConfig.Host = Host

	DbName := conf.String("mysql::mysql_db_name")
	if len(DbName) == 0 {
		logs.Error("load config of mysql_db_name failed , is null")
		err = errors.New("load config of mysql_db_name failed , is null")
		return
	}
	MysqlConfig.DbName = DbName

	return

}
