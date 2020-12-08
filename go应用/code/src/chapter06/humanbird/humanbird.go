package main

import "fmt"

// 可飞行的
type Flying struct{}

func (f *Flying) Fly() {
	fmt.Println("can fly")
}

// 可行走的
type Walkable struct{}

func (f *Walkable) Walk() {
	fmt.Println("can calk")
}

// 人类
type Human struct {
	Walkable // 人类能行走
}

// 鸟类
type Bird struct {
	Walkable // 鸟类能行走
	Flying   // 鸟类能飞行
}

func main() {

	// 实例化鸟类
	b := new(Bird)
	fmt.Println("Bird: ")
	b.Fly()
	b.Walk()

	// 实例化人类
	h := new(Human)
	fmt.Println("Human: ")
	h.Walk()

}
