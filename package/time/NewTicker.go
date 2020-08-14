/*
	Go 定时器NewTimer、NewTicker 和time.After
	定时器（time.NewTimer）

	Go语言的定时器实质是单向通道，time.Timer结构体类型中有一个time.Time类型的单向chan，源码（src/time/time.go）如下

	type Timer struct {
		C <-chan Time
		r runtimeTimer
	}

	func NewTimer
	func NewTimer(d Duration) *Timer
	NewTimer创建一个Timer，它会在最少过去时间段d后到期，向其自身的C字段发送当时的时间。
*/

// 初始化 Timer 方法为NewTimer
// 示例

package main

import (
	"time"
)

func main() {
	t := time.NewTimer(time.Second * 2)
}
