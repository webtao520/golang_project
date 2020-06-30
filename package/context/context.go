/*
	通过 sync.WaitGroup 类型优化通道对多协程协调的处理，但是现在有一个问题，
	就是我们在启动子协程之前都已经明确知道子协程的总量，如果不知道的话，该怎么实现呢？
	一种解决方案是通过 sync.WaitGroup 分批启动子协程，具体实现代码如下：
*/
/*
package main

import (
	"fmt"
	"sync"
)

func addNum(a,b  int, deferFunc func()){
    defer func(){
		  deferFunc()
	}()
	c:=a+b
	fmt.Printf("%d + %d = %d\n", a, b, c)
}

func main()  {
	total :=10
	step:=2
	fmt.Println("启动子协程")
	var  wg sync.WaitGroup
	for i:=0;i< total;i=i+step {
		wg.Add(step)
		for j:=0;j<step;j++ {
			go addNum(i+j,1,wg.Done)
		}
		wg.Wait()
	}
	fmt.Println("所有子协程执行完毕")
}
*/

/*
	这里我们采用分批次启动子协程的方法，每次通过 wg.Add() 函数设置当前批次启动的子协程数量，
	另外需要注意的是 wg.Wait() 函数最好和 wg.Add() 函数配对使用，否则可能会引起 panic。
	除此之外，我们还可以通过另一种工具实现类似需求，这就是我们今天要介绍的 context 包，这个包为我们提供了以下方法和类型：
	我们可以先通过 withXXX 方法返回一个从父 Context 拷贝的新的可撤销子 Context 对象和对应撤销函数 CancelFunc，CancelFunc 是一个函数类型，
	调用它时会撤销对应的子 Context 对象，当满足某种条件时，我们可以通过调用该函数结束所有子协程的运行，主协程在接收到信号后可以继续往后执行。
	这么说有点迷糊，下面我们结合示例代码来解释这个包的具体使用：
*/

package main

import (
    "context"
    "fmt"
    "sync/atomic"
    "time"
)

func AddNum(a *int32, b int, deferFunc func())  {
    defer func() {
        deferFunc()
    }()
    for i := 0; ; i++ {
        curNum := atomic.LoadInt32(a)
        newNum := curNum + 1
        time.Sleep(time.Millisecond * 200)
        if atomic.CompareAndSwapInt32(a, curNum, newNum) {
            fmt.Printf("number当前值: %d [%d-%d]\n", *a, b, i)
            break
        } else {
            //fmt.Printf("The CAS operation failed. [%d-%d]\n", b, i)
        }
    }
}

func main() {
    total := 10
    var num int32
    fmt.Printf("number初始值: %d\n", num)
    fmt.Println("启动子协程...")
    ctx, cancelFunc := context.WithCancel(context.Background())
    for i := 0; i < total; i++ {
        go AddNum(&num, i, func() {
            if atomic.LoadInt32(&num) == int32(total) {
                cancelFunc()
            }
        })
    }
    <- ctx.Done()
    fmt.Println("所有子协程执行完毕.")
}

/*
   在这段代码中，我们先通过 context.WithCancel 方法返回一个新的 cxt 和 cancelFunc，
   并且通过 context.Background() 方法传入父 Context，该 Context 没有值，永远不会取消，可以看作是所有 Context 的根节点，
   比如这里的 cxt 就是从父 Context 拷贝过来的可撤销的子 Context。然后我们在一个 for 循环中依次启动子协程，
   并且只有在 atomic.LoadInt32(&num) == int32(total)（所有子协程执行完毕）时调用 cancelFunc() 方法撤销对应子 Context 对象 cxt，
   这样，处于阻塞状态的 cxt.Done() 对应通道被关闭，
   我们可以接收到通道数据然后退出主程序。

   注：cxt.Done() 方法返回一个通道，该通道会在调用 cancelFunc 函数时关闭，或者在父 context 撤销时也会被关闭。

	WithDeadline 和 WithTimeout 分别比 WithCancel 多了一个 deadline 和 timeout 时间参数，表示子 Context 存活的最长时间，
	如果超过了该时间，会自动撤销对应的子 Context。
	相应的，在调用 <-cxt.Done() 等待子协程执行结束时，如果没有调用 cancelFunc 函数的话它们会等待过期时间到达自动关闭，
	不过我们通常还是会主动调用 cancelFunc 函数以便更好的控制程序运行。

	此外，context 包还提供了一个 TODO 方法，该方法用于在不知道使用哪种 Context 时使用，不过目前基本用不到，
	还有一个 withValue 方法用于返回包含上下文信息的 Context 对象，当我们需要通过 Context 传递上下文数据时可以使用该方法返回 Context：


	 ---------------------------------------------------------------------------------
	 ctx, cancelFunc := context.WithTimeout(context.Background(), 10 * time.Second)
		valueCtx := context.WithValue(ctx, "key", "value")
		defer cancelFunc()
		for i := 0; i < total; i++ {
			go AddNum(&num, i, func() {
				if atomic.LoadInt32(&num) == int32(total) {
					fmt.Println("key:", valueCtx.Value("key"))
					cancelFunc()
				}
			})
		}
		<- ctx.Done()
	 ---------------------------------------------------------------------------------
*/
