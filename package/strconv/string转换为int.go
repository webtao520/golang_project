/**
string转换为int：Atoi()

func Atoi(s string) (int, error)
由于string可能无法转换为int，所以这个函数有两个返回值：第一个返回值是转换成int的值，第二个返回值判断是否转换成功。

*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Atoi () string -> int
	i, _ := strconv.Atoi("3")
	fmt.Println(3 + i)
	fmt.Printf("%T", i)
	// Atoi() 转换失败
	i, err := strconv.Atoi("a")
	if err != nil {
		fmt.Println("converted failed")
	}
}

/*
PS D:\goLang\github\golang_project\package\strconv> go run  string转换为int.go
		6
		int
*/
