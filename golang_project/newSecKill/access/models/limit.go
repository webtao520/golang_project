package models

type TimeLimit interface {
	Count(nowTime int64) (curCount int)
	Check(nowTime int64) int
}
