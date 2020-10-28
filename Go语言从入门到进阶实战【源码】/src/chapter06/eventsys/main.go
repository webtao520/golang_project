package main

import "fmt"

// 声明角色的结构体
type Actor struct {
}

// 为角色添加一个事件处理函数
func (a *Actor) OnEvent(param interface{}) {

	fmt.Println("actor event:", param)
}

// 全局事件
func GlobalEvent(param interface{}) {

	fmt.Println("global event:", param)
}

func main() {

	// 实例化一个角色
	a := new(Actor)

	// 注册名为OnSkill的回调
	RegisterEvent("OnSkill", a.OnEvent)

	// 再次在OnSkill上注册全局事件
	RegisterEvent("OnSkill", GlobalEvent)

	// 调用事件，所有注册的同名函数都会被调用
	CallEvent("OnSkill", 100)

}
