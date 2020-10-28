package main

import (
	"bufio"
	"net"
)

// 连接的会话逻辑
func handleSession(conn net.Conn, callback func(net.Conn, []byte) bool) {

	// 创建一个socket的读取器
	dataReader := bufio.NewReader(conn)

	// 接收数据的循环
	for {

		// 从连接读取封包
		pkt, err := readPacket(dataReader)

		// 回调外部
		if err != nil || !callback(conn, pkt.Body) {

			// 回调要求退出
			conn.Close()
			break
		}
	}

}
