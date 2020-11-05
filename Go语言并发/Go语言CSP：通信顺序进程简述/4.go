package main 

import (
	"fmt"
)

func main(){
	done:=make(chan int,1) // 带缓存通道
	
	go func (){
	  fmt.Println("golang")
	  done <- 1
	}()

	data:=<-done
	fmt.Println("打印channel",data); 
}


/**
	PS D:\goLang\github\golang_project\Go语言并发\Go语言CSP：通信顺序进程简述> go run 4.go
	golang
	打印channel 1
*/