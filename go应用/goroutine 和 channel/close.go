package main

import (
	"fmt"
)

func main () {
	var intChan chan int
	intChan=make(chan int,3)
	intChan<-100
	intChan<-200 // close
	close(intChan)
	// 这是不能够再写入数到channel
	 // intChan<-300 // panic: send on closed channel
	fmt.Println("okokok!!")
	//当管道关闭后，读取数据是可以的
	n1:=<-intChan
	fmt.Println("n1=",n1)
}