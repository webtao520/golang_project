package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
)

type Blacklist struct {
	UserIdBlacklist map[int]bool
	IpBlacklist     map[string]bool
}

var (
	SecKillBlacklist *Blacklist = &Blacklist{
		UserIdBlacklist: make(map[int]bool, 10000),
		IpBlacklist:     make(map[string]bool, 10000),
	}
)

func init() {
	go LoadBlacklist()
	go SyncIpBlacklist()
	go SyncUserIdBlacklist()

	go WriteAccessRedis()
	go ReadDisposeRedis()
}

// 加载 redis 中的黑名单
func LoadBlacklist() (err error) {
	conn := BlacklistRedisPool.Get()
	reply, err := conn.Do("hgetall", "UserIdBlacklist")
	UserIdlist, err := redis.Strings(reply, err)
	if err != nil {
		logs.Warn("hget all failed, err : ", err)
		return
	}

	for _, v := range UserIdlist {
		UserId, err := strconv.Atoi(v)
		if err != nil {
			logs.Warn("invalid user UserId [%v]", UserId)
			continue
		}
		SecKillBlacklist.UserIdBlacklist[UserId] = true
	}

	reply, err = conn.Do("hgetall", "IpBlacklist")
	Iplist, err := redis.Strings(reply, err)
	if err != nil {
		logs.Warn("hget all failed, err : ", err)
		return
	}

	for _, v := range Iplist {
		SecKillBlacklist.IpBlacklist[v] = true
	}
	conn.Close()
	return
}

// 通过 redis 与 map 进行 ip黑名单 同步
func SyncIpBlacklist() {
	go SyncAddIpBlacklist()
	go SyncUnchainIpBlacklist()
}

// 处理新添加到黑名单的ip
func SyncAddIpBlacklist() {
	var IpList []string
	LastTime := time.Now().Local().Unix()
	// 定时器
	t := time.NewTicker(time.Millisecond * 100) // 100 毫秒执行一次
	for range t.C {
		conn := BlacklistRedisPool.Get()

		reply, err := conn.Do("BLPOP", "blackiplist", time.Second)
		ip, err := redis.String(reply, err)
		conn.Close()
		if err != nil {
			continue
		}
		CurTime := time.Now().Local().Unix()
		IpList = append(IpList, ip)
		if len(IpList) > 100 || CurTime-LastTime > 5 {
			RWLock.Lock()
			for _, v := range IpList {
				SecKillBlacklist.IpBlacklist[v] = true
			}
			RWLock.Unlock()

			LastTime = CurTime
			logs.Info("sync ip list from redis succ, ip[%v]", IpList)
		}

	}
}

// 处理新解除黑名单的ip
func SyncUnchainIpBlacklist() {
	// 定时器
	t := time.NewTicker(time.Millisecond * 100) // 100 毫秒执行一次
	for range t.C {
		conn := BlacklistRedisPool.Get()

		reply, err := conn.Do("BLPOP", "unblackiplist", time.Second)
		ip, err := redis.String(reply, err)
		conn.Close()
		if err != nil {
			continue
		}
		RWLock.Lock()
		if _, ok := SecKillBlacklist.IpBlacklist[ip]; ok {
			delete(SecKillBlacklist.IpBlacklist, ip)
		}
		RWLock.Unlock()
	}
}

// 通过 redis 与 map 进行 UserId黑名单 同步
func SyncUserIdBlacklist() {
	go SyncAddUserIdBlacklist()
	go SyncUnchainUserIdBlacklist()
}

func SyncAddUserIdBlacklist() {
	var UserIdList []int
	LastTime := time.Now().Local().Unix()
	// 定时器
	t := time.NewTicker(time.Millisecond * 100) // 100 毫秒执行一次
	for range t.C {
		conn := BlacklistRedisPool.Get()
		reply, err := conn.Do("BLPOP", "blackidlist", time.Second)
		UserId, err := redis.Int(reply, err)
		if err != nil {
			continue
		}

		conn.Close()
		CurTime := time.Now().Local().Unix()
		UserIdList = append(UserIdList, UserId)
		if len(UserIdList) > 100 || CurTime-LastTime > 5 {
			RWLock.Lock()
			for _, v := range UserIdList {
				SecKillBlacklist.UserIdBlacklist[v] = true
			}
			RWLock.Unlock()

		}
		logs.Info("sync id list from redis succ, UserId[%v]", UserId)
	}
}

func SyncUnchainUserIdBlacklist() {
	// 定时器
	t := time.NewTicker(time.Millisecond * 100) // 100 毫秒执行一次
	for range t.C {
		conn := BlacklistRedisPool.Get()
		reply, err := conn.Do("BLPOP", "unblackidlist", time.Second)
		userid, err := redis.Int(reply, err)
		conn.Close()
		if err != nil {
			continue
		}
		RWLock.Lock()
		if _, ok := SecKillBlacklist.UserIdBlacklist[userid]; ok {
			delete(SecKillBlacklist.UserIdBlacklist, userid)
		}
		RWLock.Unlock()
	}
}

func WriteAccessRedis() {
	for i := 0; i < SecKillConf.AccessRedisConfig.WriteGoroutineNum; i++ {
		go WriteAccessRedisNode()
	}
}

func WriteAccessRedisNode() {
	logs.Debug("WriteAccessRedisNode is running")
	for {
		req, ok := <-SecKillRequestChan
		fmt.Println("write AccessRedis : ", req)
		if !ok {
			time.Sleep(time.Millisecond * 100) // 停止100毫秒
			continue
		}
		conn := AccessRedisPool.Get()
		data, err := json.Marshal(req)
		if err != nil {
			logs.Error("json.Marshal failed, error:%v req:%v", err, req)
			conn.Close()
			continue
		}
		_, err = conn.Do("LPUSH", "request_seckill_lucky_user", string(data))
		if err != nil {
			logs.Error("lpush failed, err:%v, req:%v", err, req)
			conn.Close()
			continue
		}
		fmt.Println(string(data))
		fmt.Println("write AccessRedis sucess")
		conn.Close()
	}
}

func ReadDisposeRedis() {
	for i := 0; i < SecKillConf.AccessRedisConfig.ReadGoroutineNum; i++ {
		go ReadDisposeRedisNode()
	}
}
func ReadDisposeRedisNode() {
	logs.Debug("ReadDisposeRedisNode is running")
	for {
		conn := DisposeRedisPool.Get()
		reply, err := conn.Do("RPOP", "result_seckill_lucky_user")
		data, err := redis.String(reply, err)

		if err == redis.ErrNil {
			time.Sleep(time.Second)
			conn.Close()
			continue
		}
		logs.Debug("rpop from redis succ, data:%s", string(data))

		if err != nil {
			logs.Error("rpop failed, err:%v", err)
			conn.Close()
			continue
		}

		var result SecKillResult
		err = json.Unmarshal([]byte(data), &result)
		if err != nil {
			logs.Error("json.Unmarshal failed, err:%v", err)
			conn.Close()
			continue
		}

		UserKey := fmt.Sprintln(result.ActivityId, "seckill", result.UserId)

		MutexLock.Lock()
		resultChan, ok := SecKillUserMap[UserKey]
		MutexLock.Unlock()
		if !ok {
			conn.Close()
			logs.Warn("user not found:%v", UserKey)
			continue
		}

		resultChan <- &result

		fmt.Println("resultChan", resultChan)

		conn.Close()
	}
}
