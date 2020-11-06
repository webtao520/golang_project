package models

import (
	//"errors"
	"fmt"
	"time"
	//"math/rand"
	//"time"
)


type Antispam struct {
    UserIp  map[string]*Limit // ip 防刷
}

type Limit struct {
	SecKillSecLimit TimeLimit // 秒级
}

var (
	AntispamInfo = &Antispam{
		UserIp: make(map[int] *Limit, 10000),
	} 
)


func NewAntispamModel () *Antispam {
	return &Antispam{}
}

func (this *Antispam) SecKillAntispam(UserIp,ActivityId int,Ip string)(err error){
     if err =this.AntispamByIp(Ip);err!=nil{
		 return 
	 }
	 return 
}


// ip 防刷
func (this *Antispam)  AntispamByIp(UserIp string)(err error){
	if _, ok := SecKillBlacklist.IpBlacklist[UserIp]; ok {
		err = fmt.Errorf("ip 黑名单 ：ip : %v", UserIp)
		return
	}
	now:=time.Now().Local().Unix()
	limit, ok := AntispamInfo.UserIp[UserIp]

	if !ok {
		 limit = &Limit{
			SecKillSecLimit: &SecLimit{},
		 }
	}
	return 
}
