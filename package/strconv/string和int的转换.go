package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 最常见的是字符串和int之间的转换：
	// 1.int转换为字符串：Itoa()
	a := 5
	fmt.Println("a" + strconv.Itoa(a))
	fmt.Printf("%T", strconv.Itoa(a))
}

/*
		################# 运行结果 ###################

PS D:\goLang\github\golang_project\package\strconv> go run   string和int的转换.go
		a5
		string
*/
