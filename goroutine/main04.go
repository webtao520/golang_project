package main

import "fmt"

/*
	 数据结构：
		 数组，切片  ---> 线性表
		 栈 stack   ---> LIFO (后进先出)
		 队列queue  ----> FIFO (first in first out) 先进先出



	通道 channels
		 通道可以被认为是goroutines通信的管道，类似于管道中的水从一端到另一端的流动，数据可以从一端发送到另一端，通过通道接受

	声明通道
		 每个通道都有与其相关的类型，该类型是通道容许传输的数据类型，通道的零值为nil，nil通道没有任何用处，因此通道必须使用类似于地图和切片的方法来
		 定义

		 语言： 并发的程序之间传递数据：
		 go 语言中主张： 应该通过数据传递来实现共享内存，而不是通过共享内存来实现消息传递

		 语法：   数据类型，make() 也是引用类型的数据
			 关联一种相关的数据类型
			 nil chan,同map一样，不能使用

		操作， goroutine 可以从chan中读取数据，另一个goroutine从中写入数据
			操作符 <-
			从chan中读取数据： data := <- chan
			向chan中写入数据，chan <- data


		从chan中读取数据： 阻塞式,直到另一个goroutine 向通道中写入数据，接触阻塞
		向chan中写入数据，阻塞式，直到另一个goroutine向通道中将数据读取出，接触阻塞

		安全： 通道本身是安全的，同一时间，只能容许一个goroutine来读或写

*/

func main() {
	var ch1 chan int // 初始化
	fmt.Println(ch1)
	fmt.Printf("%T\n", ch1)

	ch1 = make(chan int) // 这样一个通道创建完成
	fmt.Println(ch1)

	ch2 := make(chan bool)

	//开启匿名函数
	go func() {
		fmt.Println("子 goroutine 。。。。")
		data := <-ch1 //阻塞式 从通道中读取数据
		fmt.Println("子goroutine 从通道中读取到main传来的数据是:", data)
		ch2 <- true // 向通道中写入数据，表示结束
	}()

	//向通道写入数据
	ch1 <- 100 //  写入 阻塞式 main goroutine 向通道中写入数据
	<-ch2
	fmt.Println("main over")
}
