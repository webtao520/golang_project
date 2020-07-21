package initall

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

var (
	Db          orm.Ormer
	SecKillConf ConfigAll
)

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
