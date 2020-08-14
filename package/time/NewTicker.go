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

	====================================================================================

	func (*Timer) Stop
	func (t *Timer) Stop() bool

	Stop停止Timer的执行。如果停止了t会返回真；如果t已经被停止或者过期了会返回假。Stop不会关闭通道t.C，
	以避免从该通道的读取不正确的成功。

*/

// 初始化 Timer 方法为NewTimer
// 示例

/*
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(time.Second * 2)
	defer t.Stop()
	for {
		<-t.C
		fmt.Println("timer running...")
		// 需要重置 Reset 使 t 重新计时
		t.Reset(time.Second * 2)
	}
}
*/

/*
	输出
	timer running…
	timer running…
	timer running…
	timer running…
	这里使用NewTimer定时器需要t.Reset重置计数时间才能接着执行。
	如果注释 t.Reset(time.Second * 2)会导致通道堵塞，报fatal error: all goroutines are asleep - deadlock!错误。
	同时需要注意 defer t.Stop()在这里并不会停止定时器。这是因为Stop会停止Timer，停止后，Timer不会再被发送，
	但是Stop不会关闭通道，防止读取通道发生错误。
*/

// 如果想停止定时器，只能让go程序自动结束

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(time.Second * 2)
	ch := make(chan bool)
	go func(t *time.Timer) {
		defer t.Stop()
		for {
			select {
			case <-t.C:
				fmt.Println("timer running....")
				// 需要重置Reset 使t 重新开始计时
				t.Reset(time.Second * 2)
			case stop := <-ch:
				if stop {
					fmt.Println("time Stop")
					return
				}
			}
		}
	}(t)
	time.Sleep(10 * time.Second)
	ch <- true
	close(ch)
	time.Sleep(1 * time.Second)
}
