package main 

import (
	"fmt"
)

func main (){
	// 管道可以声明为只读或者只写
	// 1. 再默认情况下，管道时双向的
	// var chan1 chan int // 可读可写
	// 2. 声明为只写
	var chan2 chan <- int 
	chan2 = make (chan int,3)
	chan2<-20
	// 声明为只写，假设我们来读取的话，会发生什么？
	//num:=<-chan2 // invalid operation: <-chan2 (receive from send-only type chan<- int)
	fmt.Println("chan2=",chan2)
	// 声明为只读
	var chan3 <-chan int 
	num2:=<-chan3
	fmt.Println("num2",num2)
}