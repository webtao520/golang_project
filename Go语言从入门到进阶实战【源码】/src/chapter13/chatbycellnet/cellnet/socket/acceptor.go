package socket

import (
	"chapter13/chatbycellnet/cellnet"
	"chapter13/chatbycellnet/cellnet/internal"
	"fmt"
	"net"
	"sync"
)

// 接受器
type socketAcceptor struct {
	socketPeer
	internal.SessionManager

	// 保存侦听器
	l net.Listener

	// 侦听器的停止同步
	wg sync.WaitGroup
}

// 异步开始侦听
func (a *socketAcceptor) Start(address string) cellnet.Peer {

	a.address = address

	go a.listen(address)

	return a
}

func (a *socketAcceptor) listen(address string) {

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

	log.Infof("#listen(%s) %s", a.Name(), a.address)

	// 侦听循环
	for {

		// 新连接没有到来时，Accept是阻塞的
		conn, err := a.l.Accept()

		// 发生任何的侦听错误，打印错误并退出服务器
		if err != nil {
			break
		}

		go a.onNewSession(conn)
	}
}

func (a *socketAcceptor) onNewSession(conn net.Conn) {

	ses := newSession(conn, &a.socketPeer)

	ses.start()

}

// 停止侦听器
func (a *socketAcceptor) Stop() {
	a.l.Close()
	a.wg.Wait()
}

// 实例化一个侦听器
func NewAcceptor(f cellnet.EventFunc, q cellnet.EventQueue) cellnet.Peer {
	p := &socketAcceptor{
		SessionManager: internal.NewSessionManager(),
	}

	p.queue = q
	p.eventFunc = f
	p.peerInterface = p

	return p
}
