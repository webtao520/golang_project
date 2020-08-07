package models

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Post struct {
	Id       int64
	User     *User       `orm:"rel(fk);index"`
	Title    string      `orm:"size(100);index"`
	Color    string      `orm:"size(7);index"`
	Urlname  string      `orm:"size(100);index"`
	Urltype  int8        `orm:"index"`
	Content  string      `orm:"type(text)"`
	Tags     string      `orm:"size(100);index"`
	Posttime time.Time   `orm:"auto_now_add;type(datetime);index"`
	Views    int64       `orm:"index"`
	Status   int8        `orm:"index"`
	Updated  time.Time   `orm:"auto_now_add;type(datetime);index"`
	Istop    int8        `orm:"index"`
	Cover    string      `orm:"size(70);default(/static/upload/default/blog-default-0.png)"`
	Comments []*Comments `orm:"reverse(many)"` // fk 的反向关系
}

func (m *Post) TableName() string {
	return TableName("post")
}

// 组合QuerySeter 
func (m *Post) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

// 内容URL
func (m *Post) Link() string {
	fmt.Printf("Urlname=====>%s, Urltype=====>%d\n", m.Urlname, m.Urltype)
   if m.Urlname != "" {
	   if  m.Urltype == 1{
		return fmt.Sprintf("/%s",Rawurlencode(m.Urlname))
	   }
	   return fmt.Sprintf("/article/%s",Rawurlencode(m.Urlname))
   }
   return fmt.Sprintf("/article/%d",m.Id)

}

//带链接的标签
func (m *Post) TagsLink() string {
	if m.Tags == "" {
		return ""
	}
	var buf bytes.Buffer
	arr := strings.Split(strings.Trim(m.Tags, ","), ",")
	for k, v := range arr {
		if k > 0 {
			buf.WriteString(", ")
		}
		tag := Tag{Name: v}
		if tag.Read("Name") != nil {
			return ""
		}
		buf.WriteString(tag.Link())
	}
	return buf.String()
}