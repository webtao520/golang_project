package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	// 声明一个全局的互斥锁
	// lock 是一个全局的互斥锁
	// sync 是包，synchornized同步
	// mutex 是互斥
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	// 这里我们将res放入到myMap
	// 加锁
	lock.Lock()
	myMap[n] = res
	// 解锁
	lock.Unlock()
}

func main() {
	// 开启多个协程完成这个任务[200个]
	for i := 1; i <= 200; i++ {
		go test(i)
	}
	// 休眠10秒钟 「都二个问题」
	time.Sleep(time.Second * 10)
	//输出结果
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
