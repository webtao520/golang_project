/*

什么时候需要用到锁？

当程序中就一个线程的时候，是不需要加锁的，但是通常实际的代码不会只是单线程，
所以这个时候就需要用到锁了，
那么关于锁的使用场景主要涉及到哪些呢？
	1. 多个线程在读相同的数据时
	2. 多个线程在写相同的数据时
	3. 同一个资源，有读又有写

互斥锁（sync.Mutex）

互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个 goroutine 可以访问到共享资源（同一个时刻只有一个线程能够拿到锁）
先通过一个并发读写的例子演示一下，当多线程同时访问全局变量时，结果会怎样？
*/

/**

package main

import (
	"fmt"
)

var count int

func main() {
	for i := 0; i < 2; i++ {
		go func() {
			for i := 1000000; i > 0; i-- {
				count++
			}
			fmt.Println(count)
		}()
	}

	fmt.Scanf("\n") //等待子线程全部结束
}

*/

/*---------------------------
PS D:\goLang\github\golang_project\go应用\锁\互斥锁> go run .\sync.Mutex.go
	1059036
	1062562  //最后的结果基本不可能是我们想看到的：200000
*/

//-----------------------修改代码，在累加的地方添加互斥锁，就能保证我们每次得到的结果都是想要的值

package main

import (
	"fmt"
	"sync"
)

var (
	count int
	lock  sync.Mutex
)

func main() {
	for i := 0; i < 2; i++ {
		go func() {
			for i := 1000000; i > 0; i-- {
				lock.Lock()
				count++
				lock.Unlock()
			}
			fmt.Println(count)
		}()
	}

	fmt.Scanf("\n") //等待子线程全部结束
}

/*---------------------------
PS D:\goLang\github\golang_project\go应用\锁\互斥锁> go run .\sync.Mutex.go
		1975393
		2000000
*/

/*
关于互斥锁的补充
互斥锁需要注意的问题：

不要重复锁定互斥锁
不要忘记解锁互斥锁，必要时使用 defer 语句
不要在多个函数之间直接传递互斥锁


死锁： 当前程序中的主 goroutine 以及我们启用的那些 goroutine 都已经被阻塞，
这些 goroutine 可以被称为用户级的 goroutine 这就相当于整个程序已经停滞不前了，并且这个时候 go 程序会抛出如下的 panic：

1
fatal error: all goroutines are asleep - deadlock!
并且go语言运行时系统抛出自行抛出的panic都属于致命性错误，都是无法被恢复的，调用recover函数对他们起不到任何作用

Go语言中的互斥锁是开箱即用的，也就是我们声明一个sync.Mutex 类型的变量，就可以直接使用它了，
需要注意：该类型是一个结构体类型，属于值类型的一种，将它当做参数传给一个函数，将它从函数中返回，
把它赋值给其他变量，让它进入某个管道，都会导致他的副本的产生。并且原值和副本以及多个副本之间是完全独立的，他们都是不同的互斥锁，
所以不应该将锁通过函数的参数进行传递


*/
