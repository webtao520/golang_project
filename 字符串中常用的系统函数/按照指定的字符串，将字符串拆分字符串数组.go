package main 

import (
	"fmt"
	"strings"
)

func main () {
	strArr:=strings.Split("hello,world,ok",",")
	for i:=0;i<len(strArr);i++ {
		fmt.Printf("str[%v]=%v\n",i,strArr[i])
	}
	fmt.Printf("strArr=%v\n", strArr)
}

/*
PS D:\goLang\github\golang_project\字符串中常用的系统函数> go run .\按照指定的字符串，将字符串拆分字符串数组.go
														str[0]=hello
														str[1]=world
														str[2]=ok
														strArr=[hello world ok]
*/