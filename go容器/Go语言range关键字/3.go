package main 

import (
	"fmt"
)

func main (){
	 // 创建一个整型切片，并赋值
	 slice :=[]int{10,20,30,40}
	 // 迭代每个元素。并显示其值
	 for _,value:=range slice {
		 fmt.Printf("Value: %d\n,",value)
	 }
}

/*
		PS D:\goLang\github\golang_project\go容器\Go语言range关键字> go run 3.go
		Value: 10
		,Value: 20
		,Value: 30
		,Value: 40
*/