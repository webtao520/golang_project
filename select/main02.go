package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		data := <-ch1
		fmt.Println("子goroutine中:", data)
	}()

	select {
	case ch1 <- 100:
		fmt.Println("ch1 中写出数据....")
	case ch2 <- 200:
		fmt.Println("ch2 中写出数据....")
	}
}
