package main

import (
	"net"
	"strconv"
)

func main() {

	// 测试次数
	const TestCount = 100000

	// 测试地址，如果发生侦听冲突，请调整端口
	const address = "127.0.0.1:8010"

	// 接收到的计数器
	var recvCounter int

	// 实例化一个侦听器
	acceptor := NewAcceptor()

	// 开启侦听
	acceptor.Start(address)

	// 响应封包数据
	acceptor.OnSessionData = func(conn net.Conn, data []byte) bool {

		// 转换为字符串
		str := string(data)

		// 转换为数字
		n, err := strconv.Atoi(str)

		// 任何错误或者接收错位时，报错
		if err != nil || recvCounter != n {
			panic("failed")
		}

		// 计数器增加
		recvCounter++

		// 完成计数，关闭侦听器
		if recvCounter >= TestCount {
			acceptor.Stop()
			return false
		}

		return true
	}

	// 连接器不断发送数据
	Connector(address, TestCount)

	// 等待侦听器结束
	acceptor.Wait()

}
