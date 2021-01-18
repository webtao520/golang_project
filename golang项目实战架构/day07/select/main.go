package main

import (
	"fmt"
	"sync"
)

// select 多路复用
var ch chan int
var wg sync.WaitGroup
var once sync.Once

func f1() {
	defer wg.Done()
	ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
	once.Do(func() { close(ch) }) // 确保某个操作只执行一次
}

func main() {
	wg.Add(1)
	go f1()
	wg.Wait()
}
