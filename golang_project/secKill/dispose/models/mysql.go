package models

import (
	"time"

	"github.com/astaxie/beego/logs"
)

type SecKillLuckyUser struct {
	Id         int       `orm:"pk;auto"`
	ActivityId int       `orm:"size(10);default(0)"  valid:"Required"`
	UserId     int       `orm:"size(10);default(0)"  valid:"Required"`
	Token      string    `orm:"unique;size(60);default('')" valid:"Required"`
	AddTime    time.Time `orm:"size(10);default(0)"  valid:"Min(0)"`
	Status     int       `orm:"default(1)" form:"Status" valid:"Range(1,3)"` // 1.待支付
}

func InsertLuckyUser(luckyUser *SecKillLuckyUser) (err error) {
	_, err = DB.Insert(luckyUser)
	if err != nil {
		logs.Error("insert into SecKillLuckyUser err : %v", err)
	}
	return
}
