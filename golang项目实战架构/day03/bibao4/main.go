package main

import "fmt"

// 关键字 函数名(参数)（返回值）{}
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	f1, f2 := calc(10)

	fmt.Println(f1(1), f2(2)) // 11 9
	fmt.Println(f1(3), f2(4)) // 12 8;13 9
	fmt.Println(f1(5), f2(6)) // 13 7
}
