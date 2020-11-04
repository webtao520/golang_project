package main

import (
	"fmt"
	"sync"
)

var (
	// 逻辑中使用的某个变量
	count int

	// 与变量对应的使用读写互斥锁
	countGuard sync.RWMutex
)

func GetCount() int {

	// 锁定
	countGuard.RLock()

	// 在函数退出时解除锁定
	defer countGuard.RUnlock()

	return count
}

func SetCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}

func main() {

	// 可以进行并发安全的设置
	//SetCount(1)

	// 可以进行并发安全的获取
	fmt.Println(GetCount())

}

/**
代码说明如下：
第 6 行，在声明 countGuard 时，从 sync.Mutex 互斥锁改为 sync.RWMutex 读写互斥锁。
第 12 行，获取 count 的过程是一个读取 count 数据的过程，适用于读写互斥锁。在这一行，把 countGuard.Lock() 换做 countGuard.RLock()，将读写互斥锁标记为读状态。如果此时另外一个 goroutine 并发访问了 countGuard，同时也调用了 countGuard.RLock() 时，并不会发生阻塞。
第 15 行，与读模式加锁对应的，使用读模式解锁。
*/
