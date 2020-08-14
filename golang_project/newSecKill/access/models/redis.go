package models

type  Blacklist struct {
	UserIdBlacklist map[int]bool  	// 用户黑名单
}

var (
	SecKillBlacklist  *Blacklist = &Blacklist {
		UserIdBlacklist: make(map[int]bool, 10000),
	}
)

func init() {
	 // TODO
}