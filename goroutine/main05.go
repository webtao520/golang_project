package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	go func() {
		fmt.Println("子goroutine执行。。。")
		time.Sleep(5 * time.Second)
		ch1 <- "hello"
	}()

	time.Sleep(2 * time.Second)
	data := <-ch1
	fmt.Println("main读取到的数据是：", data)
}
