package initall

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"errors"
	"github.com/astaxie/beego/logs"
)

type  ConfigAll struct {
	MysqlConfig
}

type MysqlConfig struct {
	UserName string
	PassWd   string
	Port     int
	DbName   string
	Host     string
}


var conf config.Configer
// 初始化config
func InitConfig()(ConfigAll ConfigAll,err error){
	// 加载配置文件信息 (首先初始化一个解析器对象)
	conf, err = config.NewConfig("ini", "../common/conf/common.conf")
	if err !=nil {
		err =  errors.New(fmt.Sprintf("new config failed,err:",err))
		return
	}

	MysqlConfig,err:=GetMysqlConfig()
    if err !=nil {
		return
	}
	ConfigAll.MysqlConfig = MysqlConfig
	
	return
}


func  GetMysqlConfig()(MysqlConfig MysqlConfig,err error){
	UserName := conf.String("mysql::mysql_user_name")
	if len(UserName) == 0 {
		logs.Error("load config of mysql_user_name failed, is null")
		err = errors.New("load config of mysql_user_name failed, is null")
		return
	}
	MysqlConfig.UserName=UserName

	PassWd := conf.String("mysql::mysql_pass")
	if(len(PassWd) ==0){
		logs.Error("load config of mysql_pass failed, is null")
		err = errors.New("load config of mysql_pass failed, is null")
		return
	}
	MysqlConfig.PassWd=PassWd
	
	Port, err := conf.Int("mysql::mysql_port")
	if err != nil {
		logs.Error("load config of mysql_port failed, is err : ", err)
		err = errors.New(fmt.Sprintf("load config of mysql_port failed, is err : ", err))
		return
	}
	MysqlConfig.Port=Port

	Host := conf.String("mysql::mysql_host")
	if len(Host) == 0 {
		logs.Error("load config of mysql_host failed , is null")
		err = errors.New("load config of mysql_host failed , is null")
		return
	}
	MysqlConfig.Host=Host

	DbName := conf.String("mysql::mysql_db_name")
	if len(DbName) == 0 {
		logs.Error("load config of mysql_db_name failed , is null")
		err = errors.New("load config of mysql_db_name failed , is null")
		return
	}
	MysqlConfig.DbName = DbName

	return
}