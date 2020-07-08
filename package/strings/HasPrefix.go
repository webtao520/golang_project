package main

import (
	"fmt"
	"strings"
)

func main() {
	bool := strings.HasPrefix("/go/php", "/")
	fmt.Printf("%v", bool)
}

/*
			func HasSuffix
			func HasSuffix(s, suffix string) bool
			判断s是否有后缀字符串suffix。

	PS D:\goLang\github\golang_project\package\strings> go run .\HasPrefix.go
										true
*/
