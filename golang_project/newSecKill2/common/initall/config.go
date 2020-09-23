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
	 return
}