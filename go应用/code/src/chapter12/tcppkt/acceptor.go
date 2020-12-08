package main

import (
	"fmt"
	"net"
	"sync"
)

// 接受器
type Acceptor struct {

	// 保存侦听器
	l net.Listener

	// 侦听器的停止同步
	wg sync.WaitGroup

	// 连接的数据回调
	OnSessionData func(net.Conn, []byte) bool
}

// 异步开始侦听
func (a *Acceptor) Start(address string) {

	go a.listen(address)
}

func (a *Acceptor) listen(address string) {

	// 侦听开始，添加1个任务
	a.wg.Add(1)

	// 在退出函数时，结束侦听任务
	defer a.wg.Done()

	var err error
	// 根据给定地址进行侦听
	a.l, err = net.Listen("tcp", address)

	// 如果侦听发生错误，打印错误并退出
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 侦听循环
	for {

		// 新连接没有到来时，Accept是阻塞的
		conn, err := a.l.Accept()

		// 发生任何的侦听错误，打印错误并退出服务器
		if err != nil {
			break
		}

		// 根据连接开启会话，这个过程需要在并行执行
		go handleSession(conn, a.OnSessionData)
	}
}

// 停止侦听器
func (a *Acceptor) Stop() {
	a.l.Close()
}

// 等待侦听完全停止
func (a *Acceptor) Wait() {
	a.wg.Wait()
}

// 实例化一个侦听器
func NewAcceptor() *Acceptor {
	return &Acceptor{}
}
