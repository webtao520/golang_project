/*
	func AddInt32
	func AddInt32(addr *int32, delta int32) (new int32)
	AddInt32原子性的将val的值添加到*addr并返回新值。
*/

//atomic包提供了底层的原子级内存操作
//类型共有六种：int32, int64, uint32, uint64, uintptr, unsafe.Pinter
//操作共五种：增减， 比较并交换， 载入， 存储，交换
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	//增减操作
	var a int32
	fmt.Println("a : ", a)
	//函数名以Add为前缀，加具体类型名
	//参数一，是指针类型
	//参数二，与参数一类型总是相同
	//增操作
	new_a := atomic.AddInt32(&a, 3)
	fmt.Println("new_a : ", new_a)
	//减操作
	new_a = atomic.AddInt32(&a, -2)
	fmt.Println("new_a : ", new_a)

	//CAS(Compare And Swap)比较并交换操作
	//函数名以CompareAndSwap为前缀，并具体类型名
	var b int32
	fmt.Println("b : ", b)
	//函数会先判断参数一指向的值与参数二是否相等，如果相等，则用参数三替换参数一的值。
	//最后返回是否替换成功
	atomic.CompareAndSwapInt32(&b, 0, 3)
	fmt.Println("b : ", b)

	//载入操作
	//当我们对某个变量进行读取操作时，可能该变量正在被其他操作改变，或许我们读取的是被修改了一半的数据。
	//所以我们通过Load这类函数来确保我们正确的读取
	//函数名以Load为前缀，加具体类型名
	var c int32
	wg := sync.WaitGroup{}
	//我们启100个goroutine
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tmp := atomic.LoadInt32(&c)
			if !atomic.CompareAndSwapInt32(&c, tmp, (tmp + 1)) {
				fmt.Println("c 修改失败")
			}
		}()
	}
	wg.Wait()
	//c的值有可能不等于100，频繁修改变量值情况下，CAS操作有可能不成功。
	fmt.Println("c : ", c)

	//存储操作
	//与载入函数相对应，提供原子的存储函数
	//函数名以Store为前缀，加具体类型名
	var d int32
	fmt.Println("d : ", d)
	//存储某个值时，任何CPU都不会都该值进行读或写操作
	//存储操作总会成功，它不关心旧值是什么，与CAS不同
	atomic.StoreInt32(&d, 666)
	fmt.Println("d : ", d)

	//交换操作
	//直接设置新值，返回旧值，与CAS不同，它不关心旧值。
	//函数名以Swap为前缀，加具体类型名
	var e int32
	wg2 := sync.WaitGroup{}
	//我们启10个goroutine
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			tmp := atomic.LoadInt32(&e)
			old := atomic.SwapInt32(&e, (tmp + 1))
			fmt.Println("e old : ", old)
		}()
	}
	wg2.Wait()
	fmt.Println("e : ", e)
}
