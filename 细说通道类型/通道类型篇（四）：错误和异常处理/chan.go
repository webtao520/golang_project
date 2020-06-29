/*
  引入的问题???
   在并发编程的通信过程中，最需要处理的就是超时问题：比如向通道发送数据时发现通道已满，
   或者从通道接收数据时发现通道为空。如果不正确处理这些情况，很可能会导致整个协程阻塞并产生死锁。
   此外，如果我们试图向一个已经关闭的通道发送数据或关闭已经关闭的通道，也会引发 panic。
   以上都是我们在使用通道进行并发通信时需要尤其注意的。

	接下来我们来看看如何解决上述问题


# 超时处理机制实现 #

		Go 语言没有提供直接的超时处理机制，但我们可以借助 select 语句来实现类似机制解决超时问题，
		因为 select 语句的特点是只要其中一个 case 对应的通道操作已经完成，程序就会继续往下执行，而不会考虑其他 case 的情况。
		基于此特性，我们来为通道操作实现超时处理机制，
		创建一个新的 Go 文件 channel5.go，并编写代码如下：
*/

package main

import (
	"fmt"
)

func main() {
	// 初始化 ch 通道
	ch := make(chan int, 1)

	// 初始化 timeout 通道
	//timeout := make(chan bool, 1)

	// 实现一个匿名超时等待函数
	/*
		go func() {
			time.Sleep(1e9) // 睡眠1s钟
			timeout <- true
		}()
	*/

	// 借助 timeout 通道结合 select 语句实现 ch 通道读取超时效果
	select {
	case <-ch:
		fmt.Println("接收到ch 通道数据")
		// case <-timeout:
		// 	fmt.Println("超时1s，程序退出")
	}

}

/*
	使用 select 语句可以避免永久等待的问题，因为程序会在从 timeout 通道中接收到数据后继续执行，
	无论对 ch 的读取是否还处于等待状态，从而实现 1 秒超时的效果。这种写法看起来是一个编程小技巧，
	但却是在 Go 语言并发编程中避免通道通信超时的最有效方法。
	执行上述代码，打印结果如下：

	超时1s，程序退出

	而如果没有 timeout 通道和上述 select 机制，从 ch 通道接收数据会得到如下 panic（死锁）：

	fatal error: all goroutines are asleep - deadlock!
*/
