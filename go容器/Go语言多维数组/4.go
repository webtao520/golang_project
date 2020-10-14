package main

import "fmt"

func main() {
	// 声明两个二维整型数组
	var array1 [2][2]int
	var array2 [2][2]int
	// 为array2的每个元素赋值
	array2[0][0] = 10
	array2[0][1] = 20
	array2[1][0] = 30
	array2[1][1] = 40
	array1 = array2
	// 将 array1 的索引为 1 的维度复制到一个同类型的新数组里
	var array3 [2]int = array1[1]
	fmt.Println(array3)
	// 将数组中指定的整型值赋值到新的整型变量里
	var value int = array1[1][0]
	fmt.Println(value)

}

/**
PS D:\golang\github\golang_project\go容器\Go语言多维数组> go run 4.go
[30 40]
30
*/
