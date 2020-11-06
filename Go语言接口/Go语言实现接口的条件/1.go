package main

import (
	"fmt"
)

// 定义一个数据写入器
type DataWriter   interface {
	WriteData(data interface{}) error
}

// 定义文件结构，用于实现DataWriter
type file struct {
	
}
