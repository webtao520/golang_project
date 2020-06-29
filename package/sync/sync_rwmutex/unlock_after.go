package main 

import (
	"fmt"
	"sync"
)
// unlock 使用之前不存在lock
func main (){
	var rwmutex *sync.RWMutex
	rwmutex=new (sync.RWMutex)
	rwmutex.Unlock() // panic: sync: Unlock of unlocked RWMutex
}