package models

import (
	"fmt"
)

type Antispam struct {
	 UserId   map[int]*Limit // id防刷
}

type Limit struct {
	SecKillSecLimit TimeLimit // 秒级
	SecKillMinLimit TimeLimit  // 分钟级
}

func NewAntispamModel () *Antispam {
	 return &Antispam{}
}

// 秒杀安全控制
func (this *Antispam) SecKillAntispam(UserId, ActivityId int, Ip string)(err error){
   // 用户id 防刷
   if err = this.AntispamByUserId(UserId);err !=nil {
	   return
   }
}

// UserId 防刷
func (this *Antispam) AntispamByUserId(UserId int)(err error){
    
}