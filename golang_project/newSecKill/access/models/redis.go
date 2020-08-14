package models

import (
	"time"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/logs"
)

type  Blacklist struct {
	UserIdBlacklist map[int]bool  	// 用户黑名单
}

var (
	SecKillBlacklist  *Blacklist = &Blacklist {
		UserIdBlacklist: make(map[int]bool, 10000),
	}
)

func init() {
   // go LoadBlacklist()
   //go WriteAccessRedis()
   go SyncUserIdBlacklist()
}

// 通过redis 与map 进行userId 黑名单的同步
func SyncUserIdBlacklist(){
	go SyncAddUserIdBlacklist()
}


//  加入黑名单
func SyncAddUserIdBlacklist(){
	var UserIdList []int 
	LastTime := time.Now().Local().Unix()
	// 定时器
	t:=time.NewTicker(time.Millisecond*100) // 100 ms 执行一次
	for range t.C {
		conn:=BlacklistRedisPool.Get()
		reply,err:=conn.Do("BLPOP","blackidlist",time.Second)
	}

}

