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
	// 但是我们在变量b赋值之前直接获取它的类型会发现返回的结果是nil，这看起来很奇怪，interface在Golang里也是一种类型，
	// 那它声明的变量的类型为什么是nil呢？
	fmt.Printf("%T", b)
	fmt.Println("")
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
			<nil>
			main.Sparrow
			main.Parrot
首先我们需要明确，一个接口类型变量在没有被赋值之前，它的动态类型和动态值都是 nil 。在使用 fmt.Printf("%T\n") 获取一个变量的类型时，
其实是调用了reflect包的方法进行获取的， reflect.TypeOf 获取的是接口变量的动态类型，
reflect.valueOf() 获取的是接口变量的动态值。所以 fmt.Printf("%T\n",b) 展示的是 reflect.TypeOf 的结果，
由于接口变量 b 还没有被赋值，所以它的动态类型是 nil ，动态值也会是 nil 。

对比来看，为什么只是经过了声明未赋值的变量的类型不是 nil 呢？就像在静态类型部分中所展示的那样。原因如下：
我们先来看一下 reflect.TypeOf 函数的定义，func TypeOf(i interface{}) Type{} ，
函数的参数是一个 interface 类型的变量，在调用 TypeOf 时，在接口变量 b 没有赋值之前，
它的静态类型与参数类型一致，不需要做转换，因为 b 的动态类型为 nil，所以 TypeOf 返回的结果为 nil 。
那为什么变量 i 和变量 g 的类型不为 nil 呢？当变量 i 调用 TypeOf 时，会进行类型的转换，
将int型变量i转换为 interface 型，在这个过程中会将变量 i 的类型作为 b 的动态类型，变量 i 的值（在这里是变量 i 的零值0）作为 b 的动态值。
因为 TypeOf() 获取的是变量 b 的动态类型，
所以这个时候展示出的类型为 int。

*/
