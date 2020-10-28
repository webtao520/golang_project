package main

import (
	"chapter13/chatbycellnet/cellnet"
	"chapter13/chatbycellnet/cellnet/packet"
	"chapter13/chatbycellnet/cellnet/socket"
	"chapter13/chatbycellnet/chat/proto"
	"fmt"
	"github.com/davyxu/golog"
)

var log = golog.New("main")

func onMessage(ses cellnet.Session, raw interface{}) {

	switch ev := raw.(type) {
	case socket.AcceptedEvent:
		log.Infoln("client accepted: ", ses.ID())
	case packet.MsgEvent:

		msg := ev.Msg.(*proto.ChatREQ)

		ack := proto.ChatACK{
			ID:      ses.ID(),
			Content: msg.Content,
		}

		// 广播给所有连接
		ses.Peer().VisitSession(func(ses cellnet.Session) bool {

			ses.Send(&ack)

			return true
		})

	case socket.SessionClosedEvent:
		fmt.Println("client disconnected: ", ses.ID())
	}
}

func main() {

	// 创建事件队列
	queue := cellnet.NewEventQueue()

	// 创建接受器，使用onMessage处理消息和事件
	peer := socket.NewAcceptor(packet.NewMessageCallback(onMessage), queue)

	// 开启侦听
	peer.Start("127.0.0.1:8801")

	peer.SetName("server")

	// 开启事件循环
	queue.StartLoop()

	// 等待循环退出
	queue.Wait()

	// 关闭侦听
	peer.Stop()
}
