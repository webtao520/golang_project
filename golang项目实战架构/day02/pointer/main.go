package main

import "fmt"

// pointer

func main() {
	// 1. & 取地址
	n := 18
	p := &n
	fmt.Printf("%T\n", p) // *int：int类型的指针
	fmt.Println(p)        // 0xc00000a0b0

	// 2 * 根据地址取值
	m := *p
	fmt.Println(m) // 18

	var a1 *int // nil pointer
	fmt.Println(a1)
	var a2 = new(int) // new 函数申请一个内存地址
	fmt.Println(a2)   // 0xc00000a110
	fmt.Println(*a2)
	*a2 = 100
	fmt.Println(*a2)
}

/*
PS D:\golang\github\golang_project\golang项目实战架构课程全套\day02\pointer> go run main.go
*int
0xc00000a0b0
18
<nil>
0xc00000a110
0
100
*/
