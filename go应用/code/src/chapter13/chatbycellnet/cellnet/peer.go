package cellnet

type Session interface {
	// 发送消息，消息需要以指针格式传入
	Send(msg interface{})

	// 获得原始的Socket连接
	Raw() interface{}

	// 获得Session归属的Peer
	Peer() Peer

	// 断开
	Close()

	// 标示ID
	ID() int64
}

// 事件函数的定义
type EventFunc func(interface{}) interface{}

// 发起和接受连接的通讯端
type Peer interface {

	// 开启端，传入地址
	Start(address string) Peer

	// 停止通讯端
	Stop()

	// 获取队列
	Queue() EventQueue

	// 设置事件回调函数
	SetEvent(f EventFunc)

	// 通讯端名称
	Name() string

	// 设置通讯端名称
	SetName(string)

	SessionAccessor
}

// 会话访问
type SessionAccessor interface {

	// 获取一个连接
	GetSession(int64) Session

	// 遍历连接
	VisitSession(func(Session) bool)

	// 连接数量
	SessionCount() int

	// 关闭所有连接
	CloseAllSession()
}
