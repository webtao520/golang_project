package main 

import (
	"fmt"
)

func main(){
	// 创建一个3个元素缓冲大小的整型通道
	ch:=make(chan int,3) // 缓冲大小：决定通道最多可以保存的元素数量。
	// 查看当前通道的大小
	fmt.Println(len(ch))

	// 发送3个整型元素到通道
	ch <- 1
	ch <- 2
	ch <- 3 
	// 查看当前通道的大小
	fmt.Println(len(ch))
}

/**
PS D:\goLang\github\golang_project\Go语言并发\Go语言带缓冲的通道> go run 1.go
0xc000102080
3
*/