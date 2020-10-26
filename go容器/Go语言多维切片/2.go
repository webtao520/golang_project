package main 

import (
	"fmt"
)

func main (){
// 声明一个二维整型切片并赋值
slice :=[][]int{{10,20},{100,200}}
// 为第一个切片追加值为20 的元素
fmt.Println(slice[0])
slice[0][0]=5
fmt.Println(slice)
//slice[0][0]=append(slice[0][0],5)
fmt.Println(slice)	
}


/*
PS D:\goLang\github\golang_project\go容器\Go语言多维切片> go run 2.go
[[10 20] [100 200]]
*/