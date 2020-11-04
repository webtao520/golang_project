package main 


import (
	"fmt"
	"time"
)

func main(){
	// 创建一个打点器，每 500 毫秒出发一次
	ticker:=time.NewTicker(time.Millisecond * 500)
	//fmt.Printf("%T",ticker) // *time.Tickertick

	// 创建一个计时器 ，2s后触发
	stopper := time.NewTimer(time.Second * 2)
	//fmt.Printf("%T",stopper) // *time.Timer
	
	//  声明计数变量
	var i int 
	
	//  不断地检查通道情况
	for {
		 // 多路复用通道
		 select {
		 case <-stopper.C: // 计时器到时了
		    fmt.Println("stop")
            // 跳出循环
			goto StopHere
		 case 	<-ticker.C: // 打点器触发了
			// 记录触发了多少次
			i++
			fmt.Println("tick", i)
		 }
	}

	// 退出的标签， 使用goto 跳转
	StopHere:
	fmt.Println("done")
}

/**
PS D:\goLang\github\golang_project\Go语言并发\示例：使用通道响应计时器的事件> go run 2.go
tick 1
tick 2
tick 3
tick 4
stop
done

代码说明如下：
第 11 行，创建一个打点器，500 毫秒触发一次，返回 *time.Ticker 类型变量。
第 14 行，创建一个计时器，2 秒后返回，返回 *time.Timer 类型变量。
第 17 行，声明一个变量，用于累计打点器触发次数。
第 20 行，每次触发后，select 会结束，需要使用循环再次从打点器返回的通道中获取触发通知。
第 23 行，同时等待多路计时器信号。
第 24 行，计时器信号到了。
第 29 行，通过 goto 跳出循环。
第 31 行，打点器信号到了，通过i自加记录触发次数并打印。
*/