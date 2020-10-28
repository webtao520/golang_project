package main

// 声明日志写入器接口
type LogWriter interface {
	Write(data interface{}) error
}

// 日志器
type Logger struct {

	// 这个日志器用到的日志写入器
	writerList []LogWriter
}

// 注册一个日志写入器
func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writerList = append(l.writerList, writer)
}

// 将一个data类型的数据写入到日志
func (l *Logger) Log(data interface{}) {

	// 遍历所有注册的写入器
	for _, writer := range l.writerList {

		// 将日志输出到每一个写入器
		writer.Write(data)
	}
}

// 创建日志器的实例
func NewLogger() *Logger {
	return &Logger{}
}
