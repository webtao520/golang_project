package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1)  // [5 4 3 4 5] 开头位置
	//fmt.Println(slice2) // 123
}

/*
	PS D:\golang\github\golang_project\go容器\Go语言切片复制> go run 1.go
	[1 2 3 4 5]
	[1 2 3]
*/
