package main 

import (
	"fmt"
)

func main (){
	// 声明一个2*2 的二维整形数组
	var array [2][2]int 
	// 设置每个元素的整型值
	array[0][0]=10
	array[0][1]=20
	array[1][0]=30
	array[1][1]=40
	fmt.Println(array)	
}

/**
PS D:\golang\github\golang_project\go容器\Go语言多维数组> go run 2.go
[[10 20] [30 40]]
*/