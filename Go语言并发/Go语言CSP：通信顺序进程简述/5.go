package main 

import (
	"fmt"
)

func main(){
	done:=make(chan int,10) // 带10个缓存

	// 开N个后台打印线程
	for i:=0;i<cap(done);i++{
		go func(){
			fmt.Println("golang")
			done <- 1		
		}()
	}

	// 等待N 个后台线程完成
	for i :=0;i<cap(done);i++ {
		data:=<-done
		fmt.Println("获取changel",data)
	}
}

/**
PS D:\goLang\github\golang_project\Go语言并发\Go语言CSP：通信顺序进程简述> go run 5.go
		golang
		获取changel 1
		golang
		获取changel 1
		golang
		获取changel 1
		golang
		获取changel 1
		golang
		获取changel 1
		golang
		获取changel 1
		golang
		获取changel 1
		golang
		获取changel 1
		golang
		获取changel 1
		golang
		获取changel 1
*/