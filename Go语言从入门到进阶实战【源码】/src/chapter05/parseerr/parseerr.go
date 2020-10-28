package main

import (
	"fmt"
)

// 声明一个解析错误
type ParseError struct {
	Filename string // 文件名
	Line     int    // 行号
}

// 实现error接口，返回错误描述
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.Filename, e.Line)
}

// 创建一些解析错误
func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}

func main() {

	var e error
	// 创建一个错误实例，包含文件名和行号
	e = newParseError("main.go", 1)

	// 通过error接口查看错误描述
	fmt.Println(e.Error())

	// 根据错误接口具体的类型，获取详细错误信息
	switch detail := e.(type) {
	case *ParseError: // 这是一个解析错误
		fmt.Printf("Filename: %s Line: %d\n", detail.Filename, detail.Line)
	default: // 其他类型的错误
		fmt.Println("other error")
	}
}
