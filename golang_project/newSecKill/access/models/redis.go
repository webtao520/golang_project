package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
)

type  Blacklist struct {
	IpBlacklist     map[string]bool
}

var (
	SecKillBlacklist  *Blacklist = &Blacklist{
		IpBlacklist: make(map[string]bool,10000), 
	}
)

//初始化配置信息
func init (){
	 //加载redis 中的黑名单
	 go LoadBlacklist()
}

func LoadBlacklist()(err error){
	conn := BlacklistRedisPool.Get() // 获取黑名单redis 连接

	reply, err := conn.Do("hgetall", "IpBlacklist")
	Iplist, err := redis.Strings(reply, err)
	if err != nil {
		logs.Warn("hget all failed, err : ", err)
		return
	}

	for _, v := range Iplist {
		SecKillBlacklist.IpBlacklist[v] = true
	}
	//fmt.Println("====>",SecKillBlacklist.IpBlacklist)
	conn.Close()

	return
}

