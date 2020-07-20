/*
动态类型：

接口类型变量除了静态类型之外还有动态类型，动态类型是由给接口类型变量赋值的具体值的类型来决定的，
除了动态类型之外还有动态值，动态类型和动态值是对应的，动态值就是接口类型变量赋值的具体值，
之所以被称为动态类型，是因为接口类型的动态类型是会变化的，由被赋予的值来决定。
*/

package main

import "fmt"

func main() {
	// 动态类型
	var b Bird
	// b 是main.Sparrow类型
	b = Sparrow{}
	fmt.Printf("%T\n", b)
	// b 是main.Parrot类型
	b = Parrot{}
	fmt.Printf("%T\n", b)
}

type Bird interface {
	fly()
	sing()
}

// Goose implement Bird interface
type Sparrow struct {
	age  int
	name string
}

func (s Sparrow) fly() {
	fmt.Println("I am flying.")
}
func (s Sparrow) sing() {
	fmt.Println("I can sing.")
}

type Parrot struct {
	age  int
	kind int
	name string
}

func (p Parrot) fly() {
	fmt.Println("I am flying.")
}
func (p Parrot) sing() {
	fmt.Println("I can sing.")
}

/*
PS D:\goLang\github\golang_project\go面向对象特性\golang类型> go run .\动态类型.go
			main.Sparrow
			main.Parrot
*/
