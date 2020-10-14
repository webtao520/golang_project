package main

import (
	"fmt"
)

func main () {
// 声明字符串切片
var strList []string
// 声明整型切片
var numList []int
// 声明一个空切片 已经分配了内存
var numListEmpty = []int{}
// 输出3个切片
fmt.Println(strList, numList, numListEmpty)
// 切片判定空的结果
fmt.Println(strList == nil)
fmt.Println(numList == nil)
fmt.Println(numListEmpty == nil) // numListEmpty 已经被分配到了内存，但没有元素
}