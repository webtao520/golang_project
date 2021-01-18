package main 


import (
	"fmt"
	"sync"
)

var a []int 
var b chan int // 需要指定通道中元素的类型
var wg sync.WaitGroup 

// 无缓冲通道
func  noBufChannel(){
   fmt.Println(b) // nil
   b = make(chan int) // 不带缓冲区通道的初始化
   wg.Add(1)
   go func(){
	   defer wg.Done() // goroutine - 1
	   x:=<-b  // 读chan
	   fmt.Println("后台goroutine从通道b中取到了", x)
   }()
   b <- 10 // 写chan
   fmt.Println("10发送到通道b中了...")
   wg.Wait() //  等待goroutine 计数清 0 
}

func bufChannel(){
	fmt.Println(b) // nil
	b =make(chan int,10) // 带缓冲区chan
	b<-10
	fmt.Println("10发送到通道b中了...")
	b<-20
	fmt.Println("20发送到通道b中了...")
	//读取通道中数据
	x:=<-b 
	fmt.Println("从通道b中取到了", x)
	close(b)
}

func main(){
	bufChannel()
}