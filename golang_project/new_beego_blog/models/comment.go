package models

import (
	"time"
	//"github.com/astaxie/beego/orm"
	//"math"
)

/*
			CREATE TABLE `tb_comments` (
			`id` bigint(20) NOT NULL AUTO_INCREMENT,
			`obj_pk_id` bigint(20) NOT NULL,
			`reply_pk` bigint(20) NOT NULL DEFAULT '0',
			`reply_fk` bigint(20) NOT NULL DEFAULT '0',
			`user_id` bigint(20) NOT NULL,
			`comment` longtext NOT NULL,
			`submittime` datetime NOT NULL,
			`ipaddress` varchar(255) DEFAULT NULL,
			`is_removed` tinyint(4) NOT NULL DEFAULT '0',
			`obj_pk_type` bigint(20) NOT NULL DEFAULT '0',
			PRIMARY KEY (`id`),
			KEY `tb_comments_obj_pk_id` (`obj_pk_id`),
			KEY `tb_comments_reply_pk` (`reply_pk`),
			KEY `tb_comments_reply_fk` (`reply_fk`),
			KEY `tb_comments_user_id` (`user_id`),
			KEY `tb_comments_submittime` (`submittime`),
			KEY `tb_comments_is_removed` (`is_removed`),
			KEY `tb_comments_obj_pk_type` (`obj_pk_type`)
			) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

*/

// 评论模型
type  Comments struct {
	Id          int64
	Obj_pk      *Post     `orm:"rel(fk);index"` // obj_pk_id
	Reply_pk    int64     `orm:"index"`
	Reply_fk    int64     `orm:"index"`
	User        *User     `orm:"rel(fk);index"` // user_id
	Comment     string    `orm:"type(text)"`
	Submittime  time.Time `orm:"auto_now_add;type(datetime);index"`
	Ipaddress   string    `orm:"null"`
	Is_removed  int8      `orm:"index"` //0-正常，1-删除
	Obj_pk_type int64     `orm:"index"` //0-文章评论，1-友链评论
}

// 获取表名
func (m  *Comments) TableName()string {
	return TableName("comments")
}