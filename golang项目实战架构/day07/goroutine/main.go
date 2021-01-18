package main 

import (
	"fmt"
	"time"
)

// goroutine 
func hello(){
	fmt.Println("hello")
}

// 程序启动之后会创建一个主goroutine去执行
func main(){
	   //go hello()  // 开启一个单独的goroutine去执行hello函数（任务）
	   for i :=0;i<1000;i++{
		    go func (i int) {
				fmt.Println(i) // 用的是函数参数的那个i，不是外面的i 
			}(i)
	   }
	   fmt.Println("main")
	   time.Sleep(time.Second)
	   // main函数结束了 由main函数启动的goroutine也都结束了	    
}