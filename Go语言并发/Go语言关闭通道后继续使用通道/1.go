package main

import "fmt"

func main() {
    // 创建一个整型的通道
    ch := make(chan int)

    // 关闭通道
    close(ch)

    // 打印通道的指针, 容量和长度
    fmt.Printf("ptr:%p cap:%d len:%d\n", ch, cap(ch), len(ch))

    // 给关闭的通道发送数据
    ch <- 1
}

/**
代码运行后触发宕机：
panic: send on closed channel

代码说明如下：
第 7 行，创建一个整型通道。
第 10 行，关闭通道，注意 ch 不会被 close 设置为 nil，依然可以被访问。
第 13 行，打印已经关闭通道的指针、容量和长度。
第 16 行，尝试给已经关闭的通道发送数据。

提示触发宕机的原因是给一个已经关闭的通道发送数据。
*/