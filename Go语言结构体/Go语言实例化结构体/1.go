package main

import (
	"fmt"
)

type Player struct {
	Name        string
	HealthPoint int
	MagicPoint  int
}

func main() {
	tank := new(Player) // tank 类型是 *Player  使用new 可以是结构体 整型 字符串
	tank.Name = "test"
	tank.HealthPoint = 300
	fmt.Println(tank)

}

/**
PS D:\goLang\github\golang_project\Go语言结构体\Go语言实例化结构体> go run 1.go
&{test 300 0}
*/
