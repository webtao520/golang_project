package models

type TimeLimit interface {
	Count(nowTime int64) (curCount int)
	Check(nowTime int64) int
}

// 每秒限流 防刷
type SecLimit struct {
	count   int
	curTime int64
}

func (this *SecLimit) Count(nowTime int64) (curCount int) {
	if this.curTime != nowTime {
		this.count = 1
		this.curTime = nowTime
		curCount = this.count
		return
	}

	this.count++
	curCount = this.count
	return
}

func (this *SecLimit) Check(nowTime int64) int {
	if this.curTime != nowTime {
		return 0
	}

	return this.count
}

// 每分钟限流 防刷
type MinLimit struct {
	count   int
	curTime int64
}

func (this *MinLimit) Count(nowTime int64) (curCount int) {
	if nowTime-this.curTime > 60 {
		this.count = 1
		this.curTime = nowTime
		curCount = this.count
		return
	}

	this.count++
	curCount = this.count
	return
}

func (this *MinLimit) Check(nowTime int64) int {
	if this.curTime != nowTime {
		return 0
	}

	return this.count
}

// SecondLimit 每秒限购
type SecondLimit struct {
	count   int
	curTime int64
}

func (this *SecondLimit) Count(nowTime int64) (curCount int) {
	if nowTime-this.curTime > 60 {
		this.count = 1
		this.curTime = nowTime
		curCount = this.count
		return
	}

	this.count++
	curCount = this.count
	return
}

func (this *SecondLimit) Check(nowTime int64) int {
	if this.curTime != nowTime {
		return 0
	}

	return this.count
}

// MinuteLimit 每分钟限购
type MinuteLimit struct {
	count   int
	curTime int64
}

func (this *MinuteLimit) Count(nowTime int64) (curCount int) {
	if nowTime-this.curTime > 60 {
		this.count = 1
		this.curTime = nowTime
		curCount = this.count
		return
	}

	this.count++
	curCount = this.count
	return
}

func (this *MinuteLimit) Check(nowTime int64) int {
	if this.curTime != nowTime {
		return 0
	}

	return this.count
}
