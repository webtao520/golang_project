package initall

import (
	"github.com/astaxie/beego/orm"
)

var (
	Db          orm.Ormer
	SecKillConf ConfigAll
)

/*
func InitAll() (SecKillConf ConfigAll, err error) {
	SecKillConf, err = InitConfig() // 数据库基础配置结构赋值
	fmt.Println("------->", SecKillConf)
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

/*
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
