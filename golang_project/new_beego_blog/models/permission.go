package models

import "github.com/astaxie/beego/orm"

// 权限模型
type Permission struct {
	Id   int
	Name string `orm:"unique;size(20);index"`
}

func (m *Permission) TableName() string {
	return TableName("permission")
}

func (m *Permission) Query()  orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
