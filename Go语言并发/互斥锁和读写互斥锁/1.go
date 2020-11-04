package main 

import (
	"fmt"
	"sync"
)

var (
	// 逻辑中使用的某个变量
	 count int 
	// 与变量对应的使用互斥锁
	countGuard  sync.Mutex  //  写锁
)

func GetCount () int {
	// 锁定
	countGuard.Lock()
	// 在函数退出时解除锁定
	defer countGuard.Unlock()
	return count 
}

func SetCount(c int){
	countGuard.Lock()
	count=c 
	countGuard.Unlock()
}

func main(){
	// 可以进行并发安全的设置
	SetCount(1)
	// 可以进行并发安全的获取
	fmt.Println(GetCount())
}

/**
PS D:\goLang\github\golang_project\Go语言并发\互斥锁和读写互斥锁> go run 1.go
1

代码说明如下：
第 10 行是某个逻辑步骤中使用到的变量，无论是包级的变量还是结构体成员字段，都可以。
第 13 行，一般情况下，建议将互斥锁的粒度设置得越小越好，降低因为共享访问时等待的时间。这里笔者习惯性地将互斥锁的变量命名为以下格式：
变量名+Guard

以表示这个互斥锁用于保护这个变量。
第 16 行是一个获取 count 值的函数封装，通过这个函数可以并发安全的访问变量 count。
第 19 行，尝试对 countGuard 互斥量进行加锁。一旦 countGuard 发生加锁，如果另外一个 goroutine 尝试继续加锁时将会发生阻塞，直到这个 countGuard 被解锁。
第 22 行使用 defer 将 countGuard 的解锁进行延迟调用，解锁操作将会发生在 GetCount() 函数返回时。
第 27 行在设置 count 值时，同样使用 countGuard 进行加锁、解锁操作，保证修改 count 值的过程是一个原子过程，不会发生并发访问冲突。
*/