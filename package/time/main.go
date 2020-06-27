package main

import (
	"fmt"
	"time"
)

func main() {
	/*
				  time 包对于chan的操作
					 1. Timer : 计时器
							 func NewTimer(d Duration) *Timer
							  *Timer 对象: struct : 字段C  <- chan Time

							  type Timer struct {
		    					C <-chan Time
		   						 // 内含隐藏或非导出字段
							 }

				  2. func After(d Duration) <-chan Time
	*/
	// 1. 创建计时器
	timer1 := time.NewTimer(3 * time.Second) // 创建3s之后的计时器
	fmt.Printf("%T\n", timer1)

	time1 := timer1.C
	fmt.Printf("%T\n", time1) // <-chan time.Time
	fmt.Println(time.Now())
	time2 := <-timer1.C // 等待3s获取结果  2020-06-27 14:33:43.7100408 +0800 CST m=+3.007725301
	fmt.Println(time2)

	// 2 After创建计时器 同 Timer.C
	fmt.Println("-------------------------")
	ch1 := time.After(5 * time.Second) // 创建了只读通道
	time3 := <-ch1
	fmt.Println(time3)

}
