package main

import (
	"bufio"
	"chapter13/chatbycellnet/cellnet"
	"chapter13/chatbycellnet/cellnet/packet"
	"chapter13/chatbycellnet/cellnet/socket"
	"chapter13/chatbycellnet/chat/proto"
	"fmt"
	"github.com/davyxu/golog"
	"os"
	"strings"
)

var log = golog.New("main")

// 读取控制台指令，用回调返回
func ReadConsole(callback func(string)) {

	for {

		// 从键盘读取数据，直到读到换行为止
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			break
		}

		// 去掉空白字符
		text = strings.TrimSpace(text)

		// 将数据返回给调用者
		callback(text)
	}
}

// 响应底层事件
func onMessage(ses cellnet.Session, raw interface{}) {

	switch ev := raw.(type) {
	case socket.ConnectedEvent: // 连接上服务器的事件
		fmt.Println("connected")
	case packet.MsgEvent: // 解码后的消息

		msg := ev.Msg.(*proto.ChatACK)

		log.Infof("sid%d say: %s", msg.ID, msg.Content)

	case socket.SessionClosedEvent: // 会话关闭，连接断开
		fmt.Println("disconnected ")
	}

}

func main() {

	// 创建事件队列
	queue := cellnet.NewEventQueue()

	// 创建一个连接器，传入消息处理的响应函数和事件队列
	peer := socket.NewConnector(packet.NewMessageCallback(onMessage), queue)

	// 发起连接
	peer.Start("127.0.0.1:8801")

	// 设置peer的名称
	peer.SetName("client")

	// 事件队列的循环
	queue.StartLoop()

	// 从控制台读取字符串
	ReadConsole(func(str string) {

		// 使用peer的获取会话
		ses := peer.(interface {
			Session() cellnet.Session
		}).Session()

		// 发送使用回车输入的字符串
		ses.Send(&proto.ChatREQ{
			Content: str,
		})
	})
}
