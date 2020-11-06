package   models 

type TimeLimit  interface {
	Count(nowTime int64) (curCount int)
	Check(nowTime int64) int
}

// 每秒限流  防刷
type  SecLimit strunct {
	 count int
	 curTime int64 
}



