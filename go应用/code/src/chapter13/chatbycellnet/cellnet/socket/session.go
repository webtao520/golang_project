package socket

import (
	"chapter13/chatbycellnet/cellnet"
	"chapter13/chatbycellnet/cellnet/internal"
	"io"
	"net"
	"sync"
)

// Socket会话
type socketSession struct {

	// Socket原始连接
	conn net.Conn

	// 退出同步器
	exitSync sync.WaitGroup

	// 归属的通讯端
	peer *socketPeer

	id int64

	// 发送队列
	sendChan chan interface{}
}

// 取原始连接
func (s *socketSession) Raw() interface{} {
	return s.conn
}

func (s *socketSession) ID() int64 {
	return s.id
}

func (s *socketSession) SetID(id int64) {
	s.id = id
}

func (s *socketSession) Close() {
	s.sendChan <- nil
}

// 取会话归属的通讯端
func (s *socketSession) Peer() cellnet.Peer {
	return s.peer.peerInterface
}

// 发送封包
func (s *socketSession) Send(msg interface{}) {
	s.sendChan <- msg
}

// 接收循环
func (s *socketSession) recvLoop() {

	for {

		// 发送接收消息，要求读取数据
		err := s.peer.fireEvent(RecvEvent{s})

		// 连接断开
		if err == io.EOF {
			s.peer.fireEvent(SessionClosedEvent{s, err.(error)})
			break

			// 如果由sendLoop的close造成的socket错误，conn会被置空，不会派发接收错误
		} else if err != nil && s.conn != nil {
			s.peer.fireEvent(RecvErrorEvent{s, err.(error)})
			break
		}
	}

	s.cleanup()
}

// 发送循环
func (s *socketSession) sendLoop() {

	// 遍历要发送的数据
	for msg := range s.sendChan {

		// nil表示需要退出会话通讯
		if msg == nil {
			break
		}

		// 要求发送数据
		err := s.peer.fireEvent(SendEvent{s, msg})

		// 发送错误时派发事件
		if err != nil {
			s.peer.fireEvent(SendErrorEvent{s, err.(error), msg})
			break
		}

	}

	s.cleanup()
}

// 清理资源
func (s *socketSession) cleanup() {

	// 关闭连接
	if s.conn != nil {
		s.conn.Close()
		s.conn = nil
	}

	// 关闭发送队列
	if s.sendChan != nil {
		close(s.sendChan)
		s.sendChan = nil
	}

	// 通知完成
	s.exitSync.Done()
}

// 启动会话的各种资源
func (s *socketSession) start() {

	// 将会话添加到管理器
	s.Peer().(internal.SessionManager).Add(s)

	// 会话开始工作
	s.peer.fireEvent(SessionStartEvent{s})

	// 需要接收和发送线程同时完成时才算真正的完成
	s.exitSync.Add(2)

	go func() {

		// 等待2个任务结束
		s.exitSync.Wait()

		// 在这里断开session与逻辑的所有关系
		s.peer.fireEvent(SessionExitEvent{s})

		// 将会话从管理器移除
		s.Peer().(internal.SessionManager).Remove(s)
	}()

	// 启动并发接收goroutine
	go s.recvLoop()

	// 启动并发发送goroutine
	go s.sendLoop()
}

// 默认10个长度的发送队列
const SendQueueLen = 10

func newSession(conn net.Conn, peer *socketPeer) *socketSession {
	return &socketSession{
		conn:     conn,
		peer:     peer,
		sendChan: make(chan interface{}, SendQueueLen),
	}
}
