package main

import (
	"fmt"
)

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	int
	innerS
	//在一个结构体中对于每一种数据类型只能有一个匿名字段。
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 6.6
	outer.in1 = 5
	outer.in2 = 10
	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)
	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 67, innerS{5, 10}}
	fmt.Println("outer2 is :", outer2)
}

/**
PS D:\goLang\github\golang_project\Go语言结构体\类型内嵌和结构体内嵌> go run 1.go
outer.b is: 6
outer.c is: 6.600000
outer.int is: 0
outer.in1 is: 5
outer.in2 is: 10
outer2 is : {6 7.5 67 {5 10}}
*/
