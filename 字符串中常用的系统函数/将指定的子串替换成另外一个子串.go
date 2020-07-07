package main

import (
	"fmt"
	"strings"
)

func main (){
	str:="go go hello"
	newStr:=strings.Replace(str,"go","上海",1)
	fmt.Printf("str=%v str2=%v\n",str,newStr)
	// 将指定的子串替换成另外一个子串，参数中的1 表示你希望替换几个，如果是 -1 表示全部替换
}

/*
PS D:\goLang\github\golang_project\字符串中常用的系统函数> go run .\将指定的子串替换成另外一个子串.go
														str=go go hello str2=上海 go hello
*/