package main 


import (
	"fmt"
	"sync"
)

/*
sync.WaitGroup 类型

 	sync.WaitGroup 类型是开箱即用的，也是并发安全的。该类型提供了以下三个方法：

	Add：WaitGroup 类型有一个计数器，默认值是0，我们可以通过 Add 方法来增加这个计数器的值，通常我们可以通过个方法来标记需要等待的子协程数量；

	Done：当某个子协程执行完毕后，可以通过 Done 方法标记已完成，该方法会将所属 WaitGroup 类型实例计数器值减一，通常可以通过 defer 语句来调用它；

	Wait：Wait 方法的作用是阻塞当前协程，直到对应 WaitGroup 类型实例的计数器值归零，如果在该方法被调用的时候，对应计数器的值已经是 0，那么它将不会做任何事情。
	
	至此，你可能已经看出来了，我们完全可以组合使用 sync.WaitGroup 类型提供的方法来替代之前通道中等待子协程执行完毕的实现方法，对应代码如下
*/

func add_num(a,b int,deferFunc func()){
	defer func(){
		deferFunc()
	}()
	c:=a+b
	fmt.Printf("%d + %d = %d\n",a,b,c)
}


func main (){
	 var wg sync.WaitGroup
	 wg.Add(10)
	 for i:=0;i<10;i++ {
		 go add_num(i,1,wg.Done)
	 }
	 wg.Wait()
	}

/*
看起来代码简洁多了，我们首先在主协程中声明了一个 sync.WaitGroup 类型的 wg 变量，然后调用 Add 方法设置等待子协程数为 10，然后循环启动子协程，并将 wg.Done 作为 defer 函数传递过去，最后，我们通过 wg.Wait() 等到 sync.WaitGroup 计数器值为 0 时退出程序。

上述代码打印结果和之前通过通道实现的结果是一致的：

==============================================================
0 + 1 = 1
2 + 1 = 3
9 + 1 = 10
3 + 1 = 4
4 + 1 = 5
5 + 1 = 6
6 + 1 = 7
7 + 1 = 8
8 + 1 = 9
1 + 1 = 2
==============================================================

以上就是 sync.WaitGroup 类型的典型使用场景，通过它我们可以轻松实现一主多子的协程协作。需要注意的是，该类型计数器不能小于0，否则会抛出如下 panic：

panic: sync: negative WaitGroup counter
*/