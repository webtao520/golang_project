package main

import (
	"fmt"
	"runtime"
)

// 一段耗时的计算函数
func consumer(ch chan int) {

	// 无限获取数据的循环
	for {

		// 从通道获取数据
		data := <-ch

		// 打印数据
		fmt.Println(data)
	}

}

func main() {

	// 创建一个传递数据用的通道
	ch := make(chan int)

	for {

		// 空变量，什么也不做
		var dummy string

		// 获取输入，模拟进程持续运行
		fmt.Scan(&dummy)

		// 启动并发执行consumer函数
		go consumer(ch)

		// 输出现在的goroutine数量
		fmt.Println("goroutines:", runtime.NumGoroutine())
	}

}
