package main 


import (
	"fmt"
)

/*
	for {
		// 尝试从 ch1 接收值
		data,ok:= <-ch1
	   // 尝试从ch2 接收值
	   data,ok:= <-ch2
	   ...
	}

	这种方式虽然可以实现从多个通道接收值的需求，但是运行性能会差很多。
	为了应对这种场景，Go内置了select关键字，可以同时响应多个通道的操作。


	select的使用类似于switch语句，它有一系列case分支和一个默认的分支。
	每个case会对应一个通道的通信（接收或发送）过程。select会一直等待，
	直到某个case的通信操作完成时，就会执行case分支对应的语句。具体格式如下

	select {
	case <-ch1: // 接收
	...
	case data:=<-ch2: //  接收
	...
	case ch3<-data:	 // 发送
	... 
	default:
	 默认操作	
	}

*/

// 举个例子 说明select 多路复用
func main(){
  ch:=make(chan int,1)
  for i:=0;i<10;i++{
	   select {
	   case x:=<-ch: // 接收
		  fmt.Println(x)
	   case ch<-i: // 发送
	   }
  }
}

/*
使用select语句能提高代码的可读性。
		可处理一个或多个channel的发送/接收操作。
		如果多个case同时满足，select会随机选择一个。
		对于没有case的select{}会一直等待，可用于阻塞main函数。
*/