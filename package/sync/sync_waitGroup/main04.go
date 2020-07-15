package main 
/*
与 sync.WaitGroup 类型类似，sync.Once 类型也是开箱即用和并发安全的，其主要用途是保证指定函数代码只执行一次，类似于单例模式，
常用于应用启动时的一些全局初始化操作。它只提供了一个 Do 方法，该方法只接受一个参数，
且这个参数的类型必须是 func()，即无参数无返回值的函数类型。

在具体实现时，sync.Once 还提供了一个 uint32 类型的 done 字段，它的作用是记录 Do
 传入函数被调用次数，显然，其对应的值只能是 0 和 1，之所以设置为 uint32 类型，
 是为了保证操作的原子性，回想下我们上篇教程中介绍的原子函数，再结合 Do 方法底层实现源码，即可知晓原因，这里不深入探讨了：


 func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	// Slow-path.
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

如果 done 字段的值已经是 1 了（通过 atomic.LoadUint32() 原子函数加载），表示该函数已经调用过，否则的话会调用 sync.Once 提供的互斥锁阻塞其它代码对该类型的访问，然后通过原子操作将 done 的值设置为 1，并调用传入函数。

下面我们通过一个简单的示例来演示 sync.Once 类型的使用：
*/

import (
	"fmt"
	"sync"
	"time"
)

func dosomethind(o *sync.Once) {
   fmt.Println("Start:")
    o.Do(func (){
		fmt.Println("Do Somethind....")
	})
	fmt.Println("Finished..")
   
}

func main (){
	 o:= &sync.Once{}	
	 go dosomethind(o)
	 go dosomethind(o)
	 go dosomethind(o)
	 time.Sleep(time.Second*1)
}

/*
===================================================
Start:
Do Somethind....
Finished..
Start:
Start:
Finished..
Finished..
===================================================
显然，传入 sync.Once.Do 方法的函数只会被执行一次。
*/