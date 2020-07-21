package models

import (
	"fmt"
	"newSecKill/common/initall"

	"github.com/astaxie/beego/orm"
)

var (
	DB          orm.Ormer
	SecKillConf initall.ConfigAll
)

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SecKillProduct), new(SecKillActivity))
	if err := initAll(); err != nil {
		panic(fmt.Sprintln("init database failed, err:%v", err))
	}
	// 自动建表
	orm.RunSyncdb("default", false, true)

}

func initAll() (err error) {
	if SecKillConf, err = initall.InitConfig(); err != nil { // 加载系统运行配置信息
		return
	}
	if DB, err = initall.InitMysql(); err != nil {
		return
	}
	return
}
