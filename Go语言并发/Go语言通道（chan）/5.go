package main 

import (
	"fmt"
	"time"
)

func main(){
	//  构建一个通道
	ch:=make(chan int)
	// 开启一个并发匿名函数
	go func(){
		 // 从3循环到0
		 for i:=3;i>=0;i-- {
		   // 发送3到0 之间的数值
			ch<-i 
			// 每次发送完成时等待
			time.Sleep(time.Second)	
		 }
	}()

	// 遍历接受通道数据
	for data:=range ch {
		//  打印通道数据
		fmt.Println(data)
		// 当遇到数据0时， 退出接受循环
		if data == 0 {
			break
		}
	}
}


/**
PS D:\goLang\github\golang_project\Go语言并发\Go语言通道（chan）> go run 5.go
			3
			2
			1
			0

代码说明如下：
第 12 行，通过 make 生成一个整型元素的通道。
第 15 行，将匿名函数并发执行。
第 18 行，用循环生成 3 到 0 之间的数值。
第 21 行，将 3 到 0 之间的数值依次发送到通道 ch 中。
第 24 行，每次发送后暂停 1 秒。
第 30 行，使用 for 从通道中接收数据。
第 33 行，将接收到的数据打印出来。
第 36 行，当接收到数值 0 时，停止接收。如果继续发送，由于接收 goroutine 已经退出，没有 goroutine 发送到通道，因此运行时将会触发宕机报错。
*/