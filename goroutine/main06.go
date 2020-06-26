package main

import (
	"fmt"
	"time"
)

/*
 关闭通道和通道上的范围循环
   发送者有可以通过关闭信道，来通知接收方不会有更多的数据被发送到信道上。
   接收者可以在接受来自通道的数据时使用额外的变量来检查通道是否已经关闭
   语法结构：
	   v,ok:=<- ch
	map ,存储key，value键值对
	ok的数值为true，通道正常的，读取到的data数据有效可用。
	ok的数值为false，通道关闭，读取到data是类型的零值

	通道的关闭，发送方如果数据写入完毕，可以关闭通道，用于通知接收方数据传递完毕
	   g1 ---> 写入数据
	   g2 ---> 读取数据
*/
func main() {
	ch1 := make(chan int)
	go sendData(ch1)

	//读取数据

	for {
		time.Sleep(1 * time.Second)
		data, ok := <-ch1
		if !ok {
			fmt.Println("数据读取完毕......", ok)
			break
		}
		fmt.Println("main中读取到数据 :", data)
	}
}

func sendData(ch1 chan int) {
	for i := 1; i <= 10; i++ {
		ch1 <- i
	}
	fmt.Println("发送方，写入数据完毕。。。")
	close(ch1) // 关闭通道
}
