package main

import (
	"fmt"
	"sync"
	"time"
)

/*
     RWMutex（读写锁）
			RWMutex 是单写,多读锁，该锁可以加多个读锁或者一个写锁
			读锁占用的情况下会阻止写，不会阻止读，多个 goroutine 可以同时获取读锁
			写锁会阻止其他 goroutine（无论读和写）进来，整个锁由该 goroutine 独占
			适用于读多写少的场景

     Lock() 和 Unlock()
			Lock() 加写锁，Unlock() 解写锁
			如果在加写锁之前已经有其他的读锁和写锁，则 Lock() 会阻塞直到该锁可用，为确保该锁可用，
			已经阻塞的 Lock() 调用会从获得的锁中排除新的读取器，即写锁权限高于读锁，有写锁时优先进行写锁定
			在 Lock() 之前使用 Unlock() 会导致 panic 异常

*/
func main() {
	var mutex *sync.RWMutex
	mutex = new(sync.RWMutex)
	fmt.Println("Lock the lock")
	mutex.Lock()
	fmt.Println("The lock is locked")

	channels := make([]chan int, 4)
	for i := 0; i < 4; i++ {
		channels[i] = make(chan int)
		go func(i int, c chan int) {
			fmt.Println("Not lock: ", i)
			mutex.Lock()
			fmt.Println("Locked: ", i)
			fmt.Println("Unlock the lock: ", i)
			mutex.Unlock()
			c <- i
		}(i, channels[i])
	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock")
	mutex.Unlock()
	time.Sleep(time.Second)

	for _, c := range channels {
		<-c
	}
}
