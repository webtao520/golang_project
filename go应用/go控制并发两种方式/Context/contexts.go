package main 

import (
	"context"
	"fmt"
	"time"
)

func main (){
	ctx, cancel := context.WithCancel(context.Background()) // 手动结束  context.Background() 返回一个空的Context，这个空的Context一般用于整个Context树的根节点
	go watch(ctx,"【监控1】")
    go watch(ctx,"【监控2】")
	go watch(ctx,"【监控3】")
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止了....")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string){
   for {
	   select {
	   case <-ctx.Done():
		fmt.Println(name,"监控退出了，停止了...")
		return
	   default:
		fmt.Println(name,"goroutine监控中....")
		time.Sleep(2*time.Second)
	   }
   }
}