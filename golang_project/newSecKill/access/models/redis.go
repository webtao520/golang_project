package models

import (
	"time"

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

func SyncAddUserIdBlacklist(){
	 var UserIdList []int 
	 LastTime := time.Now().Local().Unix()
	 // 定时器
	 t:=time.NewTicker(time.Millisecond * 100)  // 100毫秒执行一次
	 /*
	NewTicker返回一个新的Ticker，该Ticker包含一个通道字段，并会每隔时间段d就向该通道发送当时的时间。
	它会调整时间间隔或者丢弃tick信息以适应反应慢的接收者。如果d<=0会panic。关闭该Ticker可以释放相关资源。
	type Ticker struct {
		C <-chan Time // 周期性传递时间信息的通道
		// 内含隐藏或非导出字段
	}
	 */
	 for range t.C {
	   
	 }


}

