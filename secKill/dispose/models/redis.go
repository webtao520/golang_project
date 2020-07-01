package models

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
)

type SecKillRequest struct {
	ActivityId    int
	UserId        int
	Ip            string
	ClientAddr    string
	ClientRefence string
	CloseNotify   <-chan bool `json:"-"`

	ResultChan chan *SecKillResult `json:"-"`
}

type SecKillResult struct {
	ActivityId int
	UserId     int
	Token      string
	Error      error
}

var (
	ReadChan  chan *SecKillRequest = make(chan *SecKillRequest, 10000)
	WriteChan chan *SecKillResult  = make(chan *SecKillResult, 10000)
)

func ReadAccessRedis() {
	for i := 0; i < SecKillConf.AccessRedisConfig.ReadGoroutineNum; i++ {
		go ReadAccessRedisNode()
	}
}
func ReadAccessRedisNode() {
	logs.Debug("read goroutine running")
	for {
		conn := AccessRedisPool.Get()
		for {
			ret, err := conn.Do("blpop", "request_seckill_lucky_user", 0)
			if err != nil {
				logs.Error("pop from queue failed, err:%v", err)
				break
			}
			tmp, ok := ret.([]interface{})
			if !ok || len(tmp) != 2 {
				logs.Error("pop from queue failed, err:%v", err)
				continue
			}

			data, ok := tmp[1].([]byte)
			if !ok {
				logs.Error("pop from queue failed, err:%v", err)
				continue
			}
			logs.Debug("pop from queue, data:%s", string(data))

			var req SecKillRequest
			err = json.Unmarshal([]byte(data), &req)
			if err != nil {
				logs.Error("unmarshal to secrequest failed, err:%v", err)
				continue
			}

			// 恶意用户判断
			err = AntispamByActivityId(req)
			if err != nil {
				continue
			}

			t := time.NewTicker(time.Millisecond * time.Duration(100))
			select {
			case ReadChan <- &req:
			case <-t.C:
				logs.Warn("send to handle chan timeout, req:%v", req)
				break
			}

		}

		conn.Close()
	}
}

func Dispose() {
	for i := 0; i < 10; i++ {
		go DisposeNode()
	}
}
func DisposeNode() {
	logs.Debug("Dispose goroutine running")
	for req := range ReadChan {
		logs.Debug("begin Dispose request:%v", req)
		res, err := DisposeSecKill(req)
		if err != nil {
			logs.Warn("Dispose request %v failed, err:%v", err)
			res = &SecKillResult{
				Error: fmt.Errorf("Server busy"),
			}
		}

		timer := time.NewTicker(time.Millisecond * time.Duration(100))
		fmt.Println(res)
		select {
		case WriteChan <- res:
		case <-timer.C:
			logs.Warn("send to response chan timeout, res:%v", res)
			break
		}

	}
}

func DisposeSecKill(req *SecKillRequest) (res *SecKillResult, err error) {
	Token := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintln(time.Now().Local().Unix(), req.ActivityId, req.UserId))))
	// LuckyUser insert mysql
	LuckyUser := &SecKillLuckyUser{
		ActivityId: req.ActivityId,
		UserId:     req.UserId,
		Token:      Token,
		AddTime:    time.Now().Local(),
		Status:     1,
	}
	err = InsertLuckyUser(LuckyUser)
	res = &SecKillResult{
		ActivityId: req.ActivityId,
		UserId:     req.UserId,
		Token:      Token,
		Error:      err,
	}
	return
}

func WriteDisposeRedis() {
	for i := 0; i < SecKillConf.AccessRedisConfig.WriteGoroutineNum; i++ {
		go WriteDisposeRedisNode()
	}
}

func WriteDisposeRedisNode() {
	logs.Debug("write goroutine running")

	for res := range WriteChan {
		err := sendToRedis(res)
		if err != nil {
			logs.Error("send to redis, err:%v, res:%v", err, res)
			continue
		}
	}
}

func sendToRedis(res *SecKillResult) (err error) {

	data, err := json.Marshal(res)
	if err != nil {
		logs.Error("marshal failed, err:%v", err)
		return
	}
	fmt.Println("string data : ", string(data))
	conn := DisposeRedisPool.Get()
	_, err = conn.Do("rpush", "result_seckill_lucky_user", string(data))
	conn.Close()
	if err != nil {
		logs.Warn("rpush to redis failed, err:%v", err)
		return
	}
	fmt.Println("push result_seckill_lucky_user success")
	return
}
