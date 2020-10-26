package main 

import (
	"fmt"
)

func main (){
	// 声明一个二维切片
	//var slice [][]int 
	// 为二维切片赋值
	//slice = [][]int {{10},{100,200}}
	//fmt.Println(slice)
	// 上面的代码也可以简写为下面的样子。
	// 声明一个二维整型切片并赋值
	slice :=[][]int{{10},{100,200}}
	fmt.Println(slice)
}


/*
PS D:\goLang\github\golang_project\go容器\Go语言多维切片> go run 1.go
[[10] [100 200]]

*/