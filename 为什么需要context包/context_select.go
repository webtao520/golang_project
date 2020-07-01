package main 

import (
	"context"
	"fmt"
	"time"
)

func main (){
 ctx,cancel:=context.WithCancel(context.Background())
  // 开始goroutine ,传入ctx
  	go func (ctx context.Context) {
		  for {
			   select {
			   case <- ctx.Done():
				   fmt.Println("任务1 结束了....")
				   return
			   default:
				   fmt.Println("任务1 正在运行中.....")
				   time.Sleep(time.Second *2)
			   }
		  }
	  }(ctx)
	  
	  //运行10s后停止
	  time.Sleep(time.Second*10)
	  fmt.Println("需要停止任务1....")
	  // 使用context 的cancel 函数停止goroutine
	  cancel()
	// 为了检测监控过是否停止，如果没有监控输出，就表示停止了
	   time.Sleep(time.Second*4)
}