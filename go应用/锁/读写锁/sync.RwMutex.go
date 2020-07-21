/*
读写锁（sync.RWMutex）

在读多写少的环境中，可以优先使用读写互斥锁（sync.RWMutex），它比互斥锁更加高效。
sync 包中的 RWMutex 提供了读写互斥锁的封装
读写锁分为：读锁和写锁
如果设置了一个写锁，那么其它读的线程以及写的线程都拿不到锁，这个时候，与互斥锁的功能相同
如果设置了一个读锁，那么其它写的线程是拿不到锁的，但是其它读的线程是可以拿到锁

通过设置写锁，同样可以实现数据的一致性：

*/

package main

import (
	"fmt"
	"sync"
)

var (
	count  int
	rwLock sync.RWMutex
)

func main() {
	for i := 0; i < 2; i++ {
		go func() {
			for i := 1000000; i > 0; i-- {
				rwLock.Lock()
				count++
				rwLock.Unlock()
			}
			fmt.Println(count)
		}()
	}

	fmt.Scanf("\n") // 等待子线程全部结束
}

/*
	PS D:\goLang\github\golang_project\go应用\锁\读写锁> go run .\sync.RwMutex.go
	1855973
	2000000
*/
