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

// 系统所有配置信息
type ConfigAll struct {
	MysqlConfig
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
