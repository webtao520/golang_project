package main

import "fmt"

// 定义接口类型
type Humaner interface {
	// 方法，只有声明，没有实现，由别的类型(自定义类型)实现
	sayHi()
}

type Personer interface {
	Humaner // 匿名字段，继承了sayHi()
	sing(lrc string)
}

// 定义一个类型实现接口
type Student struct {
	name string
	id   int
}

// Student 实现了此方法
func (tmp *Student) sayHi() {
	fmt.Printf("Student[%s,%d] sayhi\n", tmp.name, tmp.id)
}

func (tmp *Student) sing(lrc string) {
	fmt.Println("Student 在唱着: ", lrc)
}

func main() {
	// 定义一个接口类型的变量
	var i Personer
	s := &Student{
		"mike",
		555,
	}
	i = s
	i.sayHi() // 继承过来的方法
	i.sing("Nihao")
}

/*
	PS D:\goLang\github\golang_project\go面向对象特性\接口> go run   .\interface03_继承.go
			Student[mike,555] sayhi
			Student 在唱着:  Nihao
*/
