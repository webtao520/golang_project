package main 

import (
	"fmt"
)

func main(){
	// 构建一个通道
	ch:=make(chan int)
	// 开启一个并发匿名函数
	go func (){
		
		fmt.Println("start goroutine")

		//通过通道通知main 的 goroutine
		ch <- 0  // 将0 写入到通道

		fmt.Println("exit goroutine")
	}()

	fmt.Println("wait goroutine")

	// 等待匿名goroutine
	<-ch   // 这个方式实际上只是通过通道在 goroutine 间阻塞收发实现并发同步。

	fmt.Println("all done")
}

/**
PS D:\goLang\github\golang_project\Go语言并发\Go语言通道（chan）> go run 4.go
wait goroutine
start goroutine
exit goroutine
all done
代码说明如下：
第 10 行，构建一个同步用的通道。
第 13 行，开启一个匿名函数的并发。
第 18 行，匿名 goroutine 即将结束时，通过通道通知 main 的 goroutine，这一句会一直阻塞直到 main 的 goroutine 接收为止。
第 27 行，开启 goroutine 后，马上通过通道等待匿名 goroutine 结束。
*/