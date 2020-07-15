package main 

import (
	"fmt"
	"time"
)
func main (){
	stop:=make(chan bool)
	go func () {
		for {
			select {
			case <-stop:  //取通道内容
				fmt.Println("监控退出了，停止了....")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2*time.Second)
			}
		}
	}()
	time.Sleep(10*time.Second)
	fmt.Println("可以了，通知监控停止...")
	stop <- true // 往通道中存内容
	 //为了检测监控过是否停止，如果没有监控输出，就表示停止了
	 time.Sleep(2 * time.Second)
}

/*
		PS D:\goLang\github\golang_project\go应用\go控制并发两种方式\WaitGroup> go run .\chan+select.go
					goroutine监控中...
					goroutine监控中...
					goroutine监控中...
					goroutine监控中...
					goroutine监控中...
					可以了，通知监控停止...
					监控退出了，停止了....
*/