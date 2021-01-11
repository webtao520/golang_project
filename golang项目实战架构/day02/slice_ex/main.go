package main

import (
	"fmt"
	"sort"
	_ "sort"
)

// 切片的练习题

func main(){
	var a = make([]int,5,10) // 创建切片，长度为5 容量为10
	fmt.Println(a) // [0 0 0 0 0]

	for i :=0;i<10;i++{
		a=append(a, i)
	}

	fmt.Println(a)  // [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
	fmt.Println(cap(a)) // 20

	var a1=[...]int{3,7,8,9,1}
	sort.Ints(a1[:]) // 对切片进行排序
	fmt.Println(a1)
}

/**
PS D:\golang\github\golang_project\golang项目实战架构课程全套\day02\slice_ex> go run .\main.go
[0 0 0 0 0]
[0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
20
[1 3 7 8 9]
*/