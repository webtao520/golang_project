/*
		sync 包还提供了一个条件变量类型 sync.Cond，它可以和互斥锁或读写锁（以下统称互斥锁）组合使用，用来协调想要访问共享资源的线程。
		不过，与互斥锁不同，条件变量 sync.Cond 的主要作用并不是保证在同一时刻仅有一个线程访问某一个共享资源，
		而是在对应的共享资源状态发送变化时，通知其它因此而阻塞的线程。条件变量总是和互斥锁组合使用，互斥锁为共享资源的访问提供互斥支持，
		而条件变量可以就共享资源的状态变化向相关线程发出通知，重在「协调」。
		下面，我们来看看如何使用条件变量 sync.Cond。
---------------------------------------------------------------------------------------
		type Cond struct {
			noCopy noCopy

			// L is held while observing or changing the condition
			L Locker

			notify  notifyList
			checker copyChecker
		}

	提供了三个方法：
	
	// 等待通知
	func (c *Cond) Wait() {
		c.checker.check()
		t := runtime_notifyListAdd(&c.notify)
		c.L.Unlock()
		runtime_notifyListWait(&c.notify, t)
		c.L.Lock()  
	}

	// 单发通知
	func (c *Cond) Signal() {
		c.checker.check()
		runtime_notifyListNotifyOne(&c.notify)  
	}

	// 广播通知
	func (c *Cond) Broadcast() {
		c.checker.check()
		runtime_notifyListNotifyAll(&c.notify)  
	}

	我们可以通过 sync.NewCond 返回对应的条件变量实例，初始化的时候需要传入互斥锁，该互斥锁实例会赋值给 sync.Cond 的 L 属性：
	locker := &sync.Mutex{}
	cond := sync.NewCond(locker)	
--------------------------------------------------------------------------------------
sync.Cond 主要实现一个条件变量，假设 goroutine A 执行前需要等待另外一个 goroutine B 的通知，那么处于等待状态的 goroutine A 会保存在一个通知列表，也就是说需要某种变量状态的 goroutine A 将会等待（Wait）在那里，当某个时刻变量状态改变时，负责通知的 goroutine B 会通过对条件变量通知的方式（Broadcast/Signal）来通知处于等待条件变量的 goroutine A，这样就可以在共享内存中实现类似「消息通知」的同步机制。

下面来看一个具体的示例。假设我们有一个读取器和一个写入器，读取器必须依赖写入器对缓冲区进行数据写入后，才可以从缓冲区中读取数据，写入器每次完成写入数据后，都需要通过某种通知机制通知处于阻塞状态的读取器，告诉它可以对数据进行访问，这种场景正好可以通过条件变量来实现：
*/

package main 

import (
	"fmt"
)

func main () {


}