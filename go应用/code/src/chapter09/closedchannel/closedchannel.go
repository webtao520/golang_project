package main

import "fmt"

func main() {
	// 创建一个整形带2个缓冲的通道
	ch := make(chan int, 2)

	// 给通道放入两个数据
	ch <- 0
	ch <- 1

	// 关闭缓冲
	close(ch)

	// 遍历缓冲所有数据，且多遍历1个
	for i := 0; i < cap(ch)+1; i++ {

		// 从通道中取出数据
		v, ok := <-ch

		// 打印取出数据的状态
		fmt.Println(v, ok)
	}
}
