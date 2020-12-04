// 取地址实例化是最广泛的一种结构体实例化方式，可以使用函数封装上面的初始化过程

package main

import (
	"fmt"
)

type Command struct {
	Name    string // 指令名称
	Var     *int   // 指令绑定的变量
	Comment string // 指令的注释
}

func newCommand(name string, varref *int, comment string) *Command {
	return &Command{
		Name:    name,
		Var:     varref,
		Comment: comment,
	}
}

var version int = 1

func main() {
	cmd := newCommand(
		"version",
		&version,
		"show version",
	)
	fmt.Println(cmd)
}

/**
PS D:\goLang\github\golang_project\Go语言结构体\Go语言实例化结构体> go run 3.go
&{version 0x563230 show version}
*/
