package main 

import (
	"fmt"
	"strings"
)

func main () {
	str1:=strings.ToLower("golang Hello")//小写
	str2:=strings.ToUpper("aa") // 大写
	fmt.Println(str1,str2)
}

/*
PS D:\goLang\github\golang_project\字符串中常用的系统函数> go run .\将字符串的字母进行大小写的转换.go
golang hello AA
*/