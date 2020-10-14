package main 

import (
	"fmt"
)

func main (){
	var team [3]string
	team[0] = "1"
	team[1]="2"
	team[2]="3"
	
	// 遍历数组
	for k,v:=range team{
		fmt.Println(k,v)
	}
}

/**
PS D:\golang\github\golang_project\go容器\Go语言数组> go run 6.go
	0 1
	1 2
	2 3

代码说明如下：
第 14 行，使用 for 循环，遍历 team 数组，遍历出的键 k 为数组的索引，值 v 为数组的每个元素值。
第 15 行，将每个键值打印出来
*/