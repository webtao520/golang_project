package test

import (
	"chapter13/chatbycellnet/cellnet"
	_ "chapter13/chatbycellnet/cellnet/codec/json"
	"chapter13/chatbycellnet/cellnet/packet"
	"chapter13/chatbycellnet/cellnet/socket"
	"fmt"
	"reflect"
	"testing"
)

type TestEchoACK struct {
	Msg   string
	Value int32
}

func server() {
	queue := cellnet.NewEventQueue()

	peer := socket.NewAcceptor(packet.NewMessageCallback(func(ses cellnet.Session, raw interface{}) {

		switch ev := raw.(type) {
		case socket.AcceptedEvent:
			fmt.Println("server accepted")
		case packet.MsgEvent:

			msg := ev.Msg.(*TestEchoACK)

			fmt.Printf("server recv %+v\n", msg)

			ses.Send(&TestEchoACK{
				Msg:   "ack",
				Value: msg.Value,
			})
		case socket.SessionClosedEvent:
			fmt.Println("server error: ", ev.Error)
		}

	}), queue)

	peer.SetName("server")

	peer.Start("127.0.0.1:8001")

	queue.StartLoop()

	queue.Wait()
}

func client() {
	queue := cellnet.NewEventQueue()

	peer := socket.NewConnector(packet.NewMessageCallback(func(ses cellnet.Session, raw interface{}) {

		switch ev := raw.(type) {
		case socket.ConnectedEvent:
			fmt.Println("client connected")
			ses.Send(&TestEchoACK{
				Msg:   "hello",
				Value: 1234,
			})
		case packet.MsgEvent:

			msg := ev.Msg.(*TestEchoACK)

			fmt.Printf("client recv %+v\n", msg)

			queue.StopLoop(0)
		case socket.SessionClosedEvent:
			fmt.Println("client error: ", ev.Error)
		}

	}), queue)

	peer.SetName("client")

	peer.Start("127.0.0.1:8001")

	queue.StartLoop()

	queue.Wait()
}

func TestOneWayEcho(t *testing.T) {

	cellnet.RegisterMessageMeta("json", // 消息的编码格式
		"test.TestEchoACK",                         // 消息名
		reflect.TypeOf((*TestEchoACK)(nil)).Elem(), // 消息的反射类型
		1, // 消息ID
	)

	go server()

	client()
}

func TestMessageMeta(t *testing.T) {

	// 注册消息
	cellnet.RegisterMessageMeta("json", // 消息的编码格式
		"test.TestEchoACK",                         // 消息名
		reflect.TypeOf((*TestEchoACK)(nil)).Elem(), // 消息的反射类型
		1, // 消息ID
	)

	// 打印消息名
	fmt.Println(cellnet.MessageMetaByID(1).Name)

	// 通过消息反射类型，查询ID
	fmt.Println(cellnet.MessageMetaByType(reflect.TypeOf(&TestEchoACK{}).Elem()).ID)

	// 使用消息名，获取消息类型
	fmt.Println(cellnet.MessageMetaByName("test.TestEchoACK").Type.Name())
}
