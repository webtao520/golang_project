package main

import "fmt"

func main() {
	/*
		// 基本数据类型在内存布局
		var i int = 10
		// i 的地址是什么 &i
		fmt.Println("i的地址=", &i)
	*/
	i := 10
	var ptr *int = &i
	fmt.Printf("ptr=%v\n, ptr = %v", ptr, &ptr)

}
