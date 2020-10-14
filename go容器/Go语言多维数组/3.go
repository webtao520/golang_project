package main 

import (
	"fmt"
)

func main (){
	// 声明两个二维整型数组
	var array1 [2][2]int 
	var array2 [2][2]int 
	// 为array2的每个元素赋值
	array2[0][0]=10
	array2[0][1]=20
	array2[1][0]=30
	array2[1][1]=40
	// 将array2的值复制给array1
	array1=array2
	fmt.Println(array1)
}

/**
PS D:\golang\github\golang_project\go容器\Go语言多维数组> go run 3.go
[[10 20] [30 40]]
*/