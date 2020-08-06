/*
	看一个题：
	查找和排序
	题目：输入任意（用户，成绩）序列，可以获得成绩从高到低或从低到高的排列,相同成绩
	都按先录入排列在前的规则处理。
	例示：
	jack 70
	peter 96
	Tom 70
	smith 67

	从高到低 成绩
	peter 96
	jack 70
	Tom 70
	smith 67

	从低到高
	smith 67
	Tom 70
	jack 70
	peter 96

	1、按照value排序
	2、可以递增排序和递减排序
	3、保证排序的稳定性


*/

// golang map按key排序

package main

import (
	"fmt"
	"sort"
)

func main() {
	//golang的map不保证有序性，所以按key排序需要取出key，对key排序，再遍历输出value

	m := make(map[int]string)
	m[1] = "a"
	m[2] = "c"
	m[3] = "b"

	// 将键按顺序存储在slice中
	var keys []int
	for k, _ := range m {
		keys = append(keys, k)
	}
	//fmt.Println(keys)
	sort.Ints(keys) // 递增顺序
	//fmt.Println(keys)

	// 进行最后排序
	for _, k := range keys {
		fmt.Println("Key:", k, "Value", m[k])
	}
}

/**
	================ 运行结果 ===================

PS D:\goLang\github\golang_project\go应用\map> go run .\map按key和按value排序.go
										Key: 1 Value a
										Key: 2 Value c
										Key: 3 Value b
*/
