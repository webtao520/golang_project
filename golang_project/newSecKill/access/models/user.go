package models

import (
	"crypto/md5"
	"fmt"
	"time"
	"github.com/astaxie/beego/logs"
	//"github.com/astaxie/beego/orm"
)


/*
type SecKillUser struct {
	UserId     int       用户id
	UserName   string    用户名
	UserPwd    string    密码
	UserEmail  string    邮箱
	UserMobile int       手机号
	Status     int       状态（1：普通用户，2：白名单，3：黑名单）
	IsLogin    int       是否登录（1：已登录，2：未登录）
	CreateTime time.Time 注册时间
}
*/
type SecKillUser struct {
	UserId     int       `orm:"pk;auto"`
	UserName   string    `orm:"unique;size(60);default('')" form:"UserName"  valid:"Required"`
	UserPwd    string    `orm:"unique;size(60);default('')" form:"UserPwd"  valid:"Required"`
	UserEmail  string    `orm:"unique;size(60);default('')" form:"UserEmail"  valid:"Email"`
	UserMobile uint64    `orm:"size(11);default(0)" form:"UserMobile" valid:"Mobile"`
	Status     uint      `orm:"default(1)" form:"Status" valid:"Range(1,3)"`
	IsLogin    uint      `orm:"default(1)" form:"IsLogin" valid:"Range(1,3)"`
	CreateTime time.Time `orm:"size(10);default(0)" form:"CreateTime" valid:"Min(0)"`
}

// 初始化
func NewUserModel() *SecKillUser {
	return &SecKillUser{}
}

// 保存用户信息
func (this *SecKillUser) InsertUser(user *SecKillUser) (num int64, err error) {
	user.UserPwd = this.UserPwdMd5(user.UserPwd)
	//fmt.Println(user)
	num, err = DB.Insert(user)
	if err != nil {
		logs.Warn("insert SecKillUser err : %v", err)
		return
	}
	return
}

// 用户密码加密
func (this *SecKillUser) UserPwdMd5(str string) string {
	data := []byte(str + "sec kill")
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}