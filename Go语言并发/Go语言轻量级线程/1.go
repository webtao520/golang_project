package main 

import (
	"fmt"
	"time"
)

func running(){
	 var times int
	 // 构建一个无限循环
	 for {
		 times ++ 
		 fmt.Println("tick",times)
		 // 延迟1s
		 time.Sleep(time.Second)
	 }
}


func main(){
	//  并发执行程序
	go running()
	// 接受命令行输入 不做任何事情
	var input string 
	fmt.Scanln(&input)
}


/**
PS D:\goLang\github\golang_project\Go语言并发\Go语言轻量级线程> go run 1.go
tick 1
tick 2
tick 3
tick 4
tick 5
tick 6
tick 7
tick 8
tick 9
*/