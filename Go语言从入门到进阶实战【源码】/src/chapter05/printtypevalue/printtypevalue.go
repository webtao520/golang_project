package main

import (
	"bytes"
	"fmt"
)

func printTypeValue(slist ...interface{}) string {

	// 字节缓冲作为快速字符串连接
	var b bytes.Buffer

	// 遍历参数
	for _, s := range slist {

		// 将interface{}类型格式化为字符串
		str := fmt.Sprintf("%v", s)

		// 类型的字符串描述
		var typeString string

		// 对s进行类型断言
		switch s.(type) {
		case bool: // 当s为布尔类型时
			typeString = "bool"
		case string: // 当s为字符串类型时
			typeString = "string"
		case int: // 当s为整形类型时
			typeString = "int"
		}

		// 写值字符串前缀
		b.WriteString("value: ")

		// 写入值
		b.WriteString(str)

		// 写类型前缀
		b.WriteString(" type: ")

		// 写类型字符串
		b.WriteString(typeString)

		// 写入换行符，一个输入变量一行
		b.WriteString("\n")

	}
	return b.String()
}

func main() {

	// 将不同类型的变量通过printTypeValue打印出来
	fmt.Println(printTypeValue(100, "str", true))
}
