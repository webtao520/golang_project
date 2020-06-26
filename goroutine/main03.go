package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
   init() 同main() 特殊的函数，由系统自动调用执行 ----> main goroutine
   runtime 包
*/

func init() {
	// 1. 获取逻辑cpu的数量
	fmt.Println("逻辑CPU的核数", runtime.NumCPU())
	// 2. 设置go程序执行的最大的: [1,256]
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	fmt.Println("hello world")
	go printNum()
	go printLetter()
	time.Sleep(3000 * time.Microsecond)
}

func printNum() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, "\t")

	}
}

func printLetter() {
	for i := 65; i < 70; i++ {
		fmt.Printf("%c\t", i)

	}
}
