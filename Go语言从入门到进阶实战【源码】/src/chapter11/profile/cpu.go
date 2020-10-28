package main

import (
	"github.com/pkg/profile"
	"time"
)

func joinSlice() []string {

	var arr []string

	for i := 0; i < 100000; i++ {
		// 故意制造很多次的切片连接，这样的操作性能较低
		arr = append(arr, "arr")
	}

	return arr
}

func main() {
	// 开始性能分析，返回一个停止接口
	stopper := profile.Start(profile.CPUProfile, profile.ProfilePath("."))

	// 在main结束时停止性能分析
	defer stopper.Stop()

	// 分析的核心逻辑
	joinSlice()

	// 让程序至少运行1秒
	time.Sleep(time.Second)
}
