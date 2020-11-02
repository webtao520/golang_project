package main 

import (
	//"fmt"
)

func main(){
	// 创建一个整型通道
	ch:=make(chan int)
	// 尝试将0通过通道发送
	ch<-0
}

/**
PS D:\goLang\github\golang_project\Go语言并发\Go语言通道（chan）> go  run 3.go
fatal error: all goroutines are asleep - deadlock!

报错的意思是：运行时发现所有的 goroutine（包括main）都处于等待 goroutine。也就是说所有 goroutine 中的 channel 并没有形成发送和接收对应的代码。
*/