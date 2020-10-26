package main 

import (
	"fmt"
)

func main (){
	// 创建一个整型切片，并赋值
	slice:=[]int{10,20,30,40}
	// 迭代每一个元素，并显示其值
	for index,value:=range slice {
		fmt.Printf("Index: %d Value: %d\n",index,value)
	}	
}


/*
PS D:\goLang\github\golang_project\go容器\Go语言range关键字> go run 1.go
Index: 0 Value: 10
Index: 1 Value: 20
Index: 2 Value: 30
Index: 3 Value: 40
*/