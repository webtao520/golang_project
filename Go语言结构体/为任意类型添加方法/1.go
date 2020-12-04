package main

import "fmt"

// 将int定义为MyInt类型
type MyInt int

// 为myInt添加IsZero（）方法
func (m MyInt) IsZero() bool {
	return m == 0
}

// 为MyInt添加Add()方法
func (m MyInt) Add(other int) int {
	return other + int(m)
}

func main() {
	var b MyInt
	fmt.Println(b.IsZero())

	b = 1

	fmt.Println(b.Add(2))
}

/**
PS D:\goLang\github\golang_project\Go语言结构体\为任意类型添加方法> go run 1.go
true
3
*/
