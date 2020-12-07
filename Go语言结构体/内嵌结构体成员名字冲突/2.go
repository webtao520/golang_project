package main

import "fmt"

type A struct {
	a int
}

type B struct {
	a int // 默认值 int 0
}

type C struct {
	A
	B
}

func main() {
	c := &C{}
	c.a = 1
	fmt.Println(c)
}

/**
PS D:\goLang\github\golang_project\Go语言结构体\内嵌结构体成员名字冲突> go run 2.go
# command-line-arguments
.\2.go:20:3: ambiguous selector c.a
*/
