package main

import (
	"fmt"
)

// 提供一个值，每次调用函数会指定对值进行累加
func Accumulate(value int) func() int {

	// 返回一个闭包
	return func() int {

		// 累加
		value++

		// 返回一个累加值
		return value
	}
}

func main() {

	// 创建一个累加器，初始值为1
	accumulator := Accumulate(1)

	// 累加1并打印
	fmt.Println(accumulator())

	fmt.Println(accumulator())

	// 打印累加器的函数地址
	fmt.Printf("%p\n", accumulator)

	// 创建一个累加器，初始值为10
	accumulator2 := Accumulate(10)

	// 累加1并打印
	fmt.Println(accumulator2())

	// 打印累加器的函数地址
	fmt.Printf("%p\n", accumulator2)
}
