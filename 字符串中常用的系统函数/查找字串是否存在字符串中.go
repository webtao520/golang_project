package main 

import (
	"fmt"
	"strings"
)

func main (){
	num:=strings.Contains("seafood","sea")  // 查找字串是否存在指定的字符串中
	fmt.Printf("num=%v\n",num)
}

/*
   PS D:\goLang\github\golang_project\字符串中常用的系统函数> go run .\查找字串是否存在字符串中.go
															num=true
*/