package main

import (
	"fmt"
)

// 定义接口类型
type Humaner interface {
	// 方法，只有声明，没有实现，由别的类型(自定义类型)实现
	sayHi()
}

type Student struct {
	name string
	id   int
}

type Teacher struct {
	addr  string
	group string
}

type Mystr string

// Teacher 实现了此方法
func (tmp *Teacher) sayHi() {
	fmt.Printf("Teacher[%s,%s] sayhi\n", tmp.addr, tmp.group)
}

// Student 实现了此方法
func (tmp *Student) sayHi() {
	fmt.Printf("Student[%s,%d] sayhi\n", tmp.name, tmp.id)
}

// Mystr 实现此方法
func (tmp *Mystr) sayHi() {
	fmt.Printf("Mystr[%s] sayHi\n", *tmp)
}

//定义一个普通函数，函数的参数为接口类型
// 只有一个函数，可以由不同变现，多态
func WhoSayHi(i Humaner) {
	fmt.Printf("%T\n", i)
	i.sayHi()
}

func main() {
	s := &Student{
		"tom",
		555,
	}
	t := &Teacher{
		"bj",
		"go",
	}
	var str Mystr = "hello golang"
	// 调用同一函数，不同表现，多态，多种形态
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

}

/*
PS D:\goLang\github\golang_project\go面向对象特性\接口> go run .\interface02_多态.go
				Student[tom,555] sayhi
				Teacher[bj,go] sayhi
				Mystr[hello golang] sayHi
*/
