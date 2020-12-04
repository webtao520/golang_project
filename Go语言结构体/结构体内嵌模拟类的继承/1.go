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


/**
PS D:\goLang\github\golang_project\Go语言结构体\结构体内嵌模拟类的继承> go run 1.go
Bird:
can fly
can calk
Human:
can calk
*/