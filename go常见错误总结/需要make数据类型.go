package main

import (
	"fmt"
)

/*
   指针，slice ，map 的零值都是nil，即还没有分配空间
   如果需要使用这样的字段，需要先make，才能使用
*/

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int // 指针  值类型都有对应的指针类型
	slice  []int
	map1   map[string]string // map
}

func main() {
	// 定义结构体变量
	var p1 Person
	fmt.Println(p1)

	if p1.ptr == nil {
		fmt.Println("ptr")
	}

	if p1.slice == nil {
		fmt.Println("slice")
	}

	if p1.map1 == nil {
		fmt.Println("map1")
	}

	// p1.slice[0]=100 // panic: runtime error: index out of range [0] with length 0
	//使用slice 再次说明一定要make
	p1.slice = make([]int, 100)
	p1.slice[0] = 100
	fmt.Println(p1.slice)

	// 使用map 一定要make
	p1.map1 = make(map[string]string)
	p1.map1["key"] = "jimmm"
	fmt.Println(p1)
}

// 不同结构体变量字段是独立，互不影响，一个结构体变量字段更改，不影响另外一个，结构体是值类型
