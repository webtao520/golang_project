package main

import (
	"fmt"
	"strings"
)

func main () {
	str:=strings.TrimSpace(" golang is very good  ")
	fmt.Printf("str=%q\n", str)
}

/*
	func TrimLeft(s string, cutset string) string   将字符串左边空格去掉
	func TrimRight(s string, cutset string) string  将字符串右边空格去掉

	PS D:\goLang\github\golang_project\字符串中常用的系统函数> go run .\将字符串左右两边的空格去掉.go
				str="golang is very good"
*/