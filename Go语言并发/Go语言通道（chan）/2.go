package main 


import (
	"fmt"
)

func main (){
	 //创建以空接口通道 
	ch:=make(chan interface{})
	 // 将0放入到通道中
	 ch<-0
	// 将hello字符串放入通道中
	ch <- "hello"
	 	 
}