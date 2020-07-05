/*
	原子性操作：
		当通过原子性方法操作某个数值时，其他的goroutine不能再访问当前的数值变量
			sync.atomic
			  针对数值：
				   操作：
						 原子加/减
						 交换
						 比较并交换：CAS
						 存储
						 加载
	非原子性操作
			 同步：
			 上锁
			 // 。。。。。

			 解锁

*/

package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var n int64
	n = 64
	fmt.Println("n的原始数值: ", n)

	// atomic.StoreInt64(&n,9)  更新数据
	// atomic.CompareAndSwapInt64(&n,3,10)   比较更新数据
	// 原子加
	newN := atomic.AddInt64(&n, 1)
	fmt.Println("新的数据：", newN)
	fmt.Println("n的新值:", n)
}
