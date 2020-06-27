package main

import (
	"fmt"
)

/*
	缓冲通道： 自带一块缓冲区，可以暂时存储数据，如果缓冲区满了，那么才会阻塞，
	非缓冲通道：默认创建的通道，都是非缓冲
*/

func main() {
	// 1 非缓冲通道
	ch1 := make(chan int)
	fmt.Println("非缓冲通道:", len(ch1), cap(ch1))
	go func() {
		<-ch1
	}()
	ch1 <- 100 //阻塞
	fmt.Println("写完了")

	// 2 缓冲通道 缓冲区数据满了，才会阻塞
	ch2 := make(chan int, 5)

	go func() {
		data1 := <-ch2
		fmt.Println("获取数据1:", data1)
	}()

	fmt.Println("缓冲通道:", len(ch2), cap(ch2))
	ch2 <- 3
	fmt.Println(len(ch2), cap(ch2))

	ch2 <- 4
	ch2 <- 5
	ch2 <- 6
	ch2 <- 7
	ch2 <- 8
	fmt.Println(len(ch2), cap(ch2)) // 缓冲区数据满了，阻塞了
}
