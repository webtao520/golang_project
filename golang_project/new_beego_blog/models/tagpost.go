package models

import (
	"time"
)

type TagPost struct {
	Id         int64
	Tag        *Tag      `orm:"rel(fk);index"`
	Postid     int64     `orm:"index"`
	Poststatus int8      `orm:"index"`
	Posttime   time.Time `auto_now_add;orm:"index"`
}

func (m *TagPost) TableName() string {
	return TableName("tag_post")
}
