// 创建一个intChan 最多可以放入3个int，演示存3数据到intChan，然后再取出这三个int
package main

import (
	"fmt"
)

func main () {
	var intChan chan int 
	intChan=make(chan int,3)
	intChan<-10
	intChan<-20
	intChan<-10
	// 因为intChan的容量为3，在存放报告 deadlock!
	//intChan<-50
   num1:=<-intChan
   num2:=<-intChan
   num3:=<-intChan
   // 因为intChan 这时已经没有数据了，再取就会报告deadlock
   // num4:=<-intChan
   fmt.Printf("num1=%v num2=%v num3=%v",num1,num2,num3)
}

/*
PS D:\goLang\github\golang_project\go应用\goroutine 和 channel> go run .\intChan.go
					num1=10 num2=20 num3=10
*/