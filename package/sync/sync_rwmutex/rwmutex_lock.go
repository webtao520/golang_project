package main

import "sync"


// RWMutex 使用不当导致的死锁
func main (){
	var rwmutex *sync.RWMutex
	rwmutex=new(sync.RWMutex) //初始化分配内存
	rwmutex.Lock()
	rwmutex.Lock()
}


// fatal error: all goroutines are asleep - deadlock!
func main01(){
    var rwmutex *sync.RWMutex
    rwmutex = new(sync.RWMutex)
    rwmutex.Lock()
    rwmutex.RLock()
}

// RUnlock() 之前不存在 RLock()
func main03(){
    var rwmutex *sync.RWMutex
    rwmutex = new(sync.RWMutex)
    rwmutex.RUnlock()   // panic: sync: RUnlock of unlocked RWMutex
}


// RUnlock() 个数多于 RLock()
func main04(){
    var rwmutex *sync.RWMutex
    rwmutex = new(sync.RWMutex)
    rwmutex.RLock() //读锁
    rwmutex.RLock()
    rwmutex.RUnlock() // 解除读锁
    rwmutex.RUnlock()
    rwmutex.RUnlock() // panic: sync: RUnlock of unlocked RWMutex
}