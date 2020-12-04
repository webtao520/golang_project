package main

import (
	"fmt"
)

// 打印消息类型，传入匿名结构体
func printMsgType(msg *struct {
	id   int
	data string
}) {
	// 使用动词 %T 打印msg的类型
	fmt.Printf("%T\n", msg)
}

func main() {
	// 实例化 一个匿名结构体
	msg := &struct { //定义部分
		id   int
		data string
	}{ //值初始化部分
		1024,
		"hello",
	}
	printMsgType(msg)
}
