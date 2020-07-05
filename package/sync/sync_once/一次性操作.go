/*
   一次性操作：
	  sync.Once
	  Do(func()) --> 设置多次，也只会执行一次
      有些操作只需要执行一次
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	once := sync.Once{}
	for i := 1; i <= 10; i++ {
		once.Do(func() {
			count++
		})
	}
	fmt.Println(count)
}
