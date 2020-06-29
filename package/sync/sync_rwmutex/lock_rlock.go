package  main 

/*
    RLock() 和 RUnlock()
			RLock() 加读锁，RUnlock() 解读锁
			RLock() 加读锁时，如果存在写锁，则无法加读锁；当只有读锁或者没有锁时，可以加读锁，读锁可以加载多个
			RUnlock() 解读锁，RUnlock() 撤销单词 RLock() 调用，对于其他同时存在的读锁则没有效果
			在没有读锁的情况下调用 RUnlock() 会导致 panic 错误
			RUnlock() 的个数不得多余 RLock()，否则会导致 panic 错误
*/

import (
	"fmt"
	"sync"
	"time"
)

// Lock() 和 RLock()
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
            fmt.Println("Not read lock: ", i)
            mutex.RLock()
            fmt.Println("Read Locked: ", i)
            fmt.Println("Unlock the read lock: ", i)
            time.Sleep(time.Second)
            mutex.RUnlock()
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