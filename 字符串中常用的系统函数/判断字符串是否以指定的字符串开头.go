package main

import (
	"fmt"
	"strings"
)

func main() {
	b := strings.HasPrefix("ftp://192.169.10.1", "ftp")
	fmt.Printf("b=%v\n", b)
}

/*
			func HasPrefix
			func HasPrefix(s, prefix string) bool
			判断s是否有前缀字符串prefix。

	PS D:\goLang\github\golang_project\字符串中常用的系统函数> go run .\判断字符串是否以指定的字符串开头.go
		b=true
*/
