/*
一个 channel 值的发送者可以 close 这个 channel，用以表示没有值会被发送了。
接收者可以通过赋值语句的第二个参数来测试 channel
是否被关闭。
当没有值可以接收并且 channel 已经被关闭，则

__

v , ok := <-ch
之后，ok会被设置为 false。
如果是

__

for i := range ch
会不断从 ch 接收值，直到 channel 被关闭。
请看示例
*/

package main

import "fmt"

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)
	for i := range ch {
		fmt.Println(i)
	}

}

/**
PS D:\goLang\github\golang_project\go常见错误总结\go语言坑之for range> go run .\range和channel的close.go
							0
							1
							1
							2
							3
							5
							8
							13
							21
							34

本例的 fibonacci(n int, ch chan int) 就是发送者。按照传送来的 n 值进行循环。循环完毕主动关闭了 ch。接收者利用 for
i := range ch 循环接收 ch 传回的每一个值。
*/
