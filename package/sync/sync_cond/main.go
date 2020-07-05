/*
		条件变量： sync.Cond,多个goroutine等待或接受通知的集合地

		L :  Locker 接口

		Cond 条件变量，总是要和锁结合使用

		3个指针方法：
		Wait()  等待goroutine 等待接受通知，Single() Broadcast() 解除阻塞
		Single() 发送通知，一个
		Broadcast() 广播 方法给所有人



	-----------------------------------------------------------
		goroutine
		   阻塞：
				   读键盘
						 ----> 键盘输入  【解除阻塞     】

					waitgroup wait()
							----> conter =0
								Add(), Done()

					chan 读写
					   ---> 写读

					cond wait()
					   ----> Single(), Broadcast()
	-----------------------------------------------------------
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	cond := sync.Cond{L: &mutex} // 条件变量
	condition := false

	go func() {
		time.Sleep(1 * time.Second)
		cond.L.Lock()
		fmt.Println("子goroutine已经锁定.......")
		fmt.Println("子goroutine更改条件数值，并发通知...")
		condition = true // 更改数值
		cond.Signal()    // 发送通知。一个goroutine,如果有多个，任意选择一个
		fmt.Println("子goroutine....继续...")
		time.Sleep(10 * time.Second)
		fmt.Println("子goroutine解锁.....")
		cond.L.Unlock()
	}()

	cond.L.Lock()
	fmt.Println("")

	if !condition {
		fmt.Println("main. 即将等待....")
		// 1. 尝试解锁
		// 2. wait 尝试解锁，解完锁之后等待其他给发通知，===> 进入了阻塞状态,等待被唤醒： single(),broadcast()
		// 3. 一旦被唤醒后，又会锁定
		cond.Wait()
		fmt.Println("main.被唤醒.....")
	}
	fmt.Println("main...继续...")
	fmt.Println("main....解锁....")
	cond.L.Unlock()
}
