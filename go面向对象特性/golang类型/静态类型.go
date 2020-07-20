/*
golang中的类型
golang中的有两种类型，静态类型（static type）和动态类型（dynamic type）

静态类型：静态类型是在声明变量的时候指定的类型，一个变量的静态类型是永恒不变的，所以才被称为静态类型，
也可以简称为类型，普通变量只有静态类型。
*/

package main

import "fmt"

type Goose struct {
	age  int
	name string
}

func main() {
	// 遍历i和变量g 只有静态类型，变量i的类型为int，g的类型为main.Goose
	var i int
	var g Goose
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", g)
}

/*
PS D:\goLang\github\golang_project\go面向对象特性\golang类型> go run .\静态类型.go
int
main.Goose
*/
