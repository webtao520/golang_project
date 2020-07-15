package main 

import (
	"context"
	"fmt"
	"time"
)

func main () {
	ctx,cancel:=context.WithCancel(context.Background())
	go func (ctx  context.Context) {
		 for {
			select {
			case <-ctx.Done():
			   fmt.Println("监控退出了，停止了....")
			   return
			default:
				fmt.Println("goroutine监控中....")
				time.Sleep(2*time.Second)
			}
		 }
	}(ctx)
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止了....")
	cancel()
	   //为了检测监控过是否停止，如果没有监控输出，就表示停止了
	 time.Sleep(5 * time.Second)
}

/*
PS D:\goLang\github\golang_project\go应用\go控制并发两种方式\Context> go run .\context.go
							goroutine监控中....
							goroutine监控中....
							goroutine监控中....
							goroutine监控中....
							goroutine监控中....
							可以了，通知监控停止了....
							监控退出了，停止了....
*/