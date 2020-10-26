package main 

import (
	"fmt"
)

func main (){
	// 创建一个整型切片，并赋值
	slice:=[]int{10,20,30,40}
	// 迭代每个元素，并显示值和地址
	for index,value:=range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice[index])
	}
}

/*
	PS D:\goLang\github\golang_project\go容器\Go语言range关键字> go run 2.go
	Value: 10 Value-Addr: C0000A0068 ElemAddr: C00009E140
	Value: 20 Value-Addr: C0000A0068 ElemAddr: C00009E148
	Value: 30 Value-Addr: C0000A0068 ElemAddr: C00009E150
	Value: 40 Value-Addr: C0000A0068 ElemAddr: C00009E158
*/