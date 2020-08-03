/*
	     range是一个内置函数。可以遍历数组、切片slice、字典map。
		当遍历数组和切片的时候，返回的是索引和元素。
		当遍历字典的时候，返回字典的键和值。
*/

package main

import (
	"fmt"
)

func main() {
	//range遍历切片的所有元素，并求和
	fibo := []int{2, 3, 5, 8, 13, 34, 55}
	sum := 0
	for _, num := range fibo {
		sum += num
		fmt.Println(num, sum)
	}
	fmt.Println("sum:", sum)
	//range遍历字符串，返回字符索引和Unicode编码。
	for i, c := range "go" {
		fmt.Printf("%d, %c\n", i, c) //%c	该值对应的unicode码值
	}
	//range遍历字典，返回键值对。
	kvs := map[string]string{"a": "apple", "b": "banana", "c": "cofox"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v) // %s	直接输出字符串或者[]byte
	}
}

/**
PS D:\goLang\github\golang_project\go常见错误总结\go语言坑之for range> go run .\range用法.go
			2 2
			3 5
			5 10
			8 18
			13 31
			34 65
			55 120
			sum: 120
			0, g
			1, o
			c -> cofox
			a -> apple
			b -> banana
*/
