package main

import (
	"fmt"
)

// 创建一个玩家生成器，输入名称，输出生成器
func playerGen(name string) func() (string, int) {

	// 血量一直为150
	hp := 150

	// 返回创建的闭包
	return func() (string, int) {

		// 将变量引用到闭包中
		return name, hp
	}
}

func main() {

	// 创建一个玩家生成器
	generator := playerGen("high noon")

	// 返回玩家的名字和血量
	name, hp := generator()

	// 打印值
	fmt.Println(name, hp)
}
