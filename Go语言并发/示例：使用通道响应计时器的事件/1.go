package main

import (
	"fmt"
	"time"
)

func main(){
	// 声明一个退出的通道
	exit:=make(chan int)
	// 打印开始
	fmt.Println("start")

	// 过1s后，调用匿名函数
	time.AfterFunc(time.Second,func(){
		   // 1s后，打印结果
	   fmt.Println("one second after")
		// 通知main() 的 goroutine 已经结束
		exit <-0
	})
	// 等待结束
	<-exit
}


/**
PS D:\goLang\github\golang_project\Go语言并发\示例：使用通道响应计时器的事件> go run 1.go
start
one second after

代码说明如下：
	第 10 行，声明一个退出用的通道，往这个通道里写数据表示退出。
	第 16 行，调用 time.AfterFunc() 函数，传入等待的时间和一个回调。回调使用一个匿名函数，在时间到达后，匿名函数会在另外一个 goroutine 中被调用。
	第 22 行，任务完成后，往退出通道中写入数值表示需要退出。
	第 26 行，运行到此处时持续阻塞，直到 1 秒后第 22 行被执行后结束阻塞。
*/