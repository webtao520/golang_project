package models

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Antispam struct {
	UserId   map[int]*Limit    // id 防刷
	UserIp   map[string]*Limit // ip 防刷
	SecSales map[int]*Limit    // 限售
}

type Limit struct {
	SecKillSecLimit TimeLimit // 秒级
	SecKillMinLimit TimeLimit // 分钟级
	SecKillSecSales TimeLimit // 秒级限购
	SecKillMinSales TimeLimit // 分钟限购
}

var (
	AntispamInfo = &Antispam{
		UserId:   make(map[int]*Limit, 10000),
		UserIp:   make(map[string]*Limit, 10000),
		SecSales: make(map[int]*Limit, 10000),
	}
	secMaxTime int = 1  // 每秒最多一次 限流防刷
	minMaxTime int = 10 // 每分最多一次 限流防刷
)

func NewAntispamModel() *Antispam {
	return &Antispam{}
}

func (this *Antispam) SecKillAntispam(UserId, ActivityId int, Ip string) (err error) {
	if err = this.AntispamByIp(Ip); err != nil {
		return
	}
	if err = this.AntispamByUserId(UserId); err != nil {
		return
	}
	if err = this.AntispamByActivityId(ActivityId); err != nil {
		return
	}
	if err = this.AntispamByUserId(ActivityId); err != nil {
		return
	}
	return
}

// ip 防刷
func (this *Antispam) AntispamByIp(UserIp string) (err error) {
	if _, ok := SecKillBlacklist.IpBlacklist[UserIp]; ok {
		err = fmt.Errorf("ip 黑名单 ：ip : %v", UserIp)
		return
	}
	now := time.Now().Local().Unix()
	limit, ok := AntispamInfo.UserIp[UserIp]
	MutexLock.Lock()
	if !ok {
		limit = &Limit{
			SecKillSecLimit: &SecLimit{},
			SecKillMinLimit: &MinLimit{},
		}
		AntispamInfo.UserIp[UserIp] = limit
	}
	SecIpCount := limit.SecKillSecLimit.Count(now)
	MinIpCount := limit.SecKillMinLimit.Count(now)
	MutexLock.Unlock()
	if SecIpCount > secMaxTime {
		err = fmt.Errorf("ip 秒级防刷超过次数")
		return
	}
	if MinIpCount > minMaxTime {
		err = fmt.Errorf("ip 分钟级防刷超过次数")
		return
	}
	fmt.Println(AntispamInfo.UserIp[UserIp].SecKillMinLimit, "~~~", AntispamInfo.UserIp[UserIp].SecKillSecLimit)
	return
}

// UserId 防刷
func (this *Antispam) AntispamByUserId(UserId int) (err error) {
	if _, ok := SecKillBlacklist.UserIdBlacklist[UserId]; ok {
		err = fmt.Errorf("UserId 黑名单用户，UserId : %v", UserId)
		return
	}
	now := time.Now().Local().Unix()
	limit, ok := AntispamInfo.UserId[UserId]
	MutexLock.Lock()
	if !ok {
		limit = &Limit{
			SecKillSecLimit: &SecLimit{},
			SecKillMinLimit: &MinLimit{},
		}
		AntispamInfo.UserId[UserId] = limit
	}
	SecIpCount := limit.SecKillSecLimit.Count(now)
	MinIpCount := limit.SecKillMinLimit.Count(now)
	MutexLock.Unlock()
	fmt.Println(AntispamInfo)

	if SecIpCount > secMaxTime {
		err = fmt.Errorf("UserId 秒级防刷超过次数")
		return
	}
	if MinIpCount > minMaxTime {
		err = fmt.Errorf("UserId 分钟级防刷超过次数")
		return
	}
	return
}

// ActivityId 判断
func (this *Antispam) AntispamByActivityId(ActivityId int) (err error) {
	Activity, ok := SecKillInfoMap[ActivityId]
	if !ok {
		err = fmt.Errorf("ActivityId : %d 没找到 Activity 信息", ActivityId)
		fmt.Println(err)
		return
	}
	if Activity.Status != 2 {
		err = fmt.Errorf("ActivityId : %d ,Activity.Status != 2", ActivityId)
		fmt.Println(err)
		return
	}
	now := time.Now().Local()
	if !now.After(Activity.StartTime) {
		err = fmt.Errorf("Activity.StartTime 异常 , StartTime : %v , NOW : %v", Activity.StartTime, now)
		fmt.Println(err)
		return
	}
	if !Activity.EndTime.After(now) {
		err = fmt.Errorf("Activity.EndTime 异常 , EndTime : %v , NOW : %v", Activity.EndTime, now)
		fmt.Println(err)
		return
	}

	// 判断每秒可售数量
	limit, ok := AntispamInfo.SecSales[Activity.ActivityId]
	MutexLock.Lock()
	if !ok {
		limit = &Limit{
			SecKillSecSales: &SecondLimit{},
		}
		AntispamInfo.SecSales[Activity.ActivityId] = limit
	}
	SecSalesCount := limit.SecKillSecSales.Count(now.Unix())
	MutexLock.Unlock()
	if SecSalesCount > Activity.SecondLimit {
		err = fmt.Errorf("ActivityId : %d 秒限购 now : %v", ActivityId, now)
		return
	}
	// fmt.Println(AntispamInfo)
	// fmt.Println(Activity)

	return
}

// 抢到概率判断
func (this *Antispam) AntispamByBuyRate(ActivityId int) (err error) {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano()) //利用当前时间的UNIX时间戳初始化rand包
	randNumber := rand.Intn(100)
	if randNumber > SecKillInfoMap[ActivityId].BuyRate {
		err = errors.New("大于秒杀中奖概率")
		return
	}
	return
}
