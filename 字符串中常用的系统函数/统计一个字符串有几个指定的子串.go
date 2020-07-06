package main 



import (
	"fmt"
	"strings"
)

func main () {
  num:=strings.Count("ceheese","e")
  // 查找一个字符串有几个指定的子串
  fmt.Printf("num=%v\n",num)
}
/*
-------------------执行结果--------------------------
PS D:\goLang\github\golang_project\字符串中常用的系统函数> go run .\统计一个字符串有几个指定的子串.go
															num=4

*/