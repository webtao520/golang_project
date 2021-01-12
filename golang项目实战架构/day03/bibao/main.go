package main

import "fmt"

// 闭包

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

// 要求：
// f1(f2)

func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}

func main() {
	ret := f3(f2, 100, 200) // 把原来需要传递两个int类型的参数包装成一个不需要传参的函数
	fmt.Printf("%T\n", ret)
	// ret()
	f1(ret)
}


/**
func()
this is f1
this is f2
300
*/