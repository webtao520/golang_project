package models

import (
	"fmt"
	"newSecKill/common/initall"

	"github.com/astaxie/beego/orm"
	"github.com/beego/admin/src/models"
)

var (
	DB          orm.Ormer
	SecKillConf initall.ConfigAll
)

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SecKillProduct), new(SecKillActivity))
	//数据库连接
	models.Connect()
	// 自动建表
	orm.RunSyncdb("default", false, true)

	DB = orm.NewOrm()

	if err := initAll(); err != nil {
		panic(fmt.Sprintln("init database failed, err:%v", err))
	}

}

func initAll() (err error) {
	if SecKillConf, err = initall.InitConfig(); err != nil {
		return
	}
	// fmt.Println("-------->", SecKillConf)  // --------> {{root root123 3306 admin_admin localhost}}
	if err != nil {
		return
	}
	return
}
