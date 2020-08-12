package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// 标签表
type Tag struct {
	Id      int64
	Name    string     `orm:"size(20);index"`
	Count   int64      `orm:"index"`
	TagPost []*TagPost `orm:"reverse(many)"` // 多
}

func (m *Tag) TableName() string {
	return TableName("tag")
}

// 读取
func (m *Tag) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

//标签连接
func (m *Tag) Link() string {
	return fmt.Sprintf("<a class=\"category\" href=\"/category/%d\">%s</a>", m.Id, m.Name)
}

//表查询
func (m *Tag) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
