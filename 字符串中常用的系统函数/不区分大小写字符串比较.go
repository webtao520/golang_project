package main 

import (
	"fmt"
	"strings"
)

func main () {
   b:=strings.EqualFold("abc","AbC")
   fmt.Printf("b=%v\n",b)
   // == 区分大小写比较
   fmt.Println("结果","abc" == "Abc")  // 结果 false
}

/*
	------------------------执行结果-------------------------
	PS D:\goLang\github\golang_project\字符串中常用的系统函数> go run .\不区分大小写字符串比较.go
															 b=true
*/