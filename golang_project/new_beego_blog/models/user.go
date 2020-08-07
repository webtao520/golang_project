package models

import (
	//"fmt"
	"time"
	//"github.com/astaxie/beego/orm"
)

// 用户表模型
type User struct {
	Id         int64
	Username   string    `orm:"unique;size(15);index"`
	Password   string    `orm:"size(32)"`
	Nickname   string    `orm:"size(15);index"`
	Email      string    `orm:"size(50);index"`
	Lastlogin  time.Time `orm:"auto_now_add;type(datetime);index"`
	Logincount int64     `orm:"index"`
	Lastip     string    `orm:"size(32);index"`
	Authkey    string    `orm:"size(10)"`
	Active     int8
	Permission string `orm:"size(100);index"`
	Avator     string `orm:"size(150);default(/static/upload/default/user-default-60x60.png)"`
	Upcount    int64
	Post       []*Post     `orm:"reverse(many)"`
	Comments   []*Comments `orm:"reverse(many)"`
}

// 获取表名
func (m *User) TableName() string {
	//fmt.Println("自定义表名====>",TableName("user"))  自定义表名====> tb_user
	return TableName("user")
}
