package main

import (
	"fmt"
	"sync"
	"time"
)

/*
		 互斥锁： true/false 锁头对象
		 互斥： 锁定，解锁
		 有两个指针方法:
			lock() 上锁
			      阻塞的：goroutine上锁，其他的goroutine处于阻塞状态
			unlock() 开锁
*/
func main(){
	  var mutex sync.Mutex
	  fmt.Println("main 即将锁定mutex ....")
	  mutex.Lock()
	  fmt.Println("main 已经锁定mutex")
	  for i:=1;i<=3;i++ {
		  go func (i int){
			   fmt.Println("子goroutine",i,"即将锁定mutex...")
			   mutex.Lock() //阻塞
			   fmt.Println("子goroutine",i,"已经锁定...")
		  }(i)
		 time.Sleep(5*time.Second)
		 fmt.Println("main即将解锁")
		 mutex.Unlock() 
		 fmt.Println("main已经解锁")
		// mutex.Unlock()  // fatal error: sync: unlock of unlocked mutex
		 time.Sleep(3*time.Second)
	  }
}