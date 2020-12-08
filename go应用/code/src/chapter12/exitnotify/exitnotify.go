package main

import (
	"fmt"
	"net"
	"time"
)

// 套接字接收过程
func socketRecv(conn net.Conn, exitChan chan string) {

	// 创建一个接收的缓冲
	buff := make([]byte, 1024)

	// 不停的接收数据
	for {

		// 从套接字中读取数据
		_, err := conn.Read(buff)

		// 需要结束接收，退出循环
		if err != nil {
			break
		}

	}

	// 函数已经结束，发送通知
	exitChan <- "recv exit"
}

func main() {

	// 连接一个地址
	conn, err := net.Dial("tcp", "www.163.com:80")

	// 发生错误时打印错误退出
	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建退出通道
	exit := make(chan string)

	// 并发执行套接字接收
	go socketRecv(conn, exit)

	// 在接收时，等待1秒
	time.Sleep(time.Second)

	// 主动关闭套接字
	conn.Close()

	// 等待goroutine退出完毕
	fmt.Println(<-exit)
}
