package main

import (
	"errors"
	"fmt"
	"os"
)

// 文件写入器
type fileWriter struct {
	file *os.File
}

// 设置文件写入器写入的文件名
func (f *fileWriter) SetFile(filename string) (err error) {

	// 如果文件已经打开，关闭前一个文件
	if f.file != nil {
		f.file.Close()
	}

	// 创建一个文件并保存文件句柄
	f.file, err = os.Create(filename)

	// 如果创建的过程出现错误，则返回错误
	return err
}

// 实现LogWriter的Write
func (f *fileWriter) Write(data interface{}) error {

	// 日志文件可能没有创建成功
	if f.file == nil {

		// 日志文件没有准备好
		return errors.New("file not created")
	}

	// 将数据序列化为字符串
	str := fmt.Sprintf("%v\n", data)

	// 将数据以字节数组写入到文件
	_, err := f.file.Write([]byte(str))

	return err
}

// 创建文件写入器实例
func newFileWriter() *fileWriter {
	return &fileWriter{}
}
