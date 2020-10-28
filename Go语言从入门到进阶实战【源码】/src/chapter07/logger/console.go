package main

import (
	"fmt"
	"os"
)

// 命令行写入器
type consoleWriter struct {
}

// 实现LogWriter的Write
func (f *consoleWriter) Write(data interface{}) error {

	// 将数据序列化为字符串
	str := fmt.Sprintf("%v\n", data)

	// 将数据以字节数组写入到命令行
	_, err := os.Stdout.Write([]byte(str))

	return err
}

// 创建命令行写入器实例
func newConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}
