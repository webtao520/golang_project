/*
golang中的类型断言
因为接口变量的动态类型是变化的，有时我们需要知道一个接口变量的动态类型究竟是什么，这就需要使用类型断言，
断言就是对接口变量的类型进行检查，其语法结构如下：

value, ok := x.(T)
x表示要断言的接口变量；
T表示要断言的目标类型；
value表示断言成功之后目标类型变量；
ok表示断言的结果，是一个bool型变量，true表示断言成功，false表示失败，如果失败value的值为nil。
*/

package main

import "fmt"

func main() {
	// 动态类型
	var b Bird
	fmt.Printf("%T", b)
	// b 是 main.Sparrow 类型
	b = Sparrow{}
	Listen(b)

}

func Listen(b Bird) {
	if _, ok := b.(Sparrow); ok {
		fmt.Println("Listren sparrow sing.")
	}

	if _, ok := b.(Parrot); ok {
		fmt.Println("Listren parrot sing.")
	}

	b.sing()
}

type Bird interface {
	fly()
	sing()
}

type Sparrow struct {
	age  int
	name string
}

func (s Sparrow) fly() {
	fmt.Println("I am sparrow, i can fly.")
}
func (s Sparrow) sing() {
	fmt.Println("I am sparrow, i can sing.")
}

type Parrot struct {
	age  int
	kind int
	name string
}

func (p Parrot) fly() {
	fmt.Println("I am parrot, i can fly.")
}
func (p Parrot) sing() {
	fmt.Println("I am parrot, i can sing.")
}

/*
PS D:\goLang\github\golang_project\go面向对象特性\golang类型> go run .\类型断言.go
<nil>Listren sparrow sing.
I am sparrow, i can sing.
*/
