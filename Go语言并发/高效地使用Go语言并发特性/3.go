// 建议使用等待组来实现同步
package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	// 连接一个地址
	conn, err := net.Dial("tcp", "www.163.com:80")

	// 发生错误时打印错误退出
	if err != nil {
		fmt.Println(err)
		return
	}

	// 退出通道
	var wg sync.WaitGroup
	// 添加一个任务
	wg.Add(1)

	// 并发执行接受套接字
	go socketRecv(conn, &wg)

	// 在接受时，等待1s
	time.Sleep(time.Second)

	// 主动关闭套接字
	conn.Close()
	// 等待goroutine退出完毕
	wg.Wait()
	fmt.Println("recv done")
}

func socketRecv(conn net.Conn, wg *sync.WaitGroup) {
	// 创建一个接受的缓冲
	buff := make([]byte, 1024)

	// 不停地接受数据
	for {
		// 从套接字中读取数据
		_, err := conn.Read(buff)

		// 需要结束接受，退出循环
		if err != nil {
			break
		}
	}
	// 函数已经结束，发送通知
	wg.Done()
}
