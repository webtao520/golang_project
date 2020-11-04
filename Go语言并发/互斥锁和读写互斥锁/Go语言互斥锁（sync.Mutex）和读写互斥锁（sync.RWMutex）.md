Go语言包中的 sync 包提供了两种锁类型：sync.Mutex 和 sync.RWMutex。

Mutex 是最简单的一种锁类型，同时也比较暴力，当一个 goroutine 获得了 Mutex 后，
其他 goroutine 就只能乖乖等到这个 goroutine 释放该 Mutex。

RWMutex 相对友好些，是经典的单写多读模型。在读锁占用的情况下，会阻止写，但不阻止读，也就是多个 goroutine 可同时获取读锁
（调用 RLock() 方法；而写锁（调用 Lock() 方法）会阻止任何其他 goroutine（无论读和写）进来，整个锁相当于由该 goroutine 独占。
从 RWMutex 的实现看，RWMutex 类型其实组合了 Mutex：

type RWMutex struct {
    w Mutex
    writerSem uint32
    readerSem uint32
    readerCount int32
    readerWait int32
}

对于这两种锁类型，任何一个 Lock() 或 RLock() 均需要保证对应有 Unlock() 或 RUnlock() 调用与之对应，否则可能导致等待该锁的所有 goroutine 处于饥饿状态，
甚至可能导致死锁。锁的典型使用模式如下：
+ 案例
 * 1.go 

 在读多写少的环境中，可以优先使用读写互斥锁（sync.RWMutex），它比互斥锁更加高效。sync 包中的 RWMutex 提供了读写互斥锁的封装。

我们将互斥锁例子中的一部分代码修改为读写互斥锁，参见下面代码：
 + 案例
    * 2.go 