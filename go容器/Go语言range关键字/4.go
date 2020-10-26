package main 

import (
	"fmt"
)

func main (){
	 // 创建一个整型切片，并赋值
	 slice:=[]int{10,20,30,40}
	 // 从第三个元素开始迭代每个元素
	 for index:=2;index <len(slice);index++ {
		fmt.Printf("Index: %d Value: %d\n", index, slice[index])
	 }
}

/*
	PS D:\goLang\github\golang_project\go容器\Go语言range关键字> go run 4.go
	Index: 2 Value: 30
	Index: 3 Value: 40
*/