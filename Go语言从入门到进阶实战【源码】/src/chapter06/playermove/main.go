package main

import "fmt"

func main() {

	// 实例化玩家对象，并设速度为0.5
	p := NewPlayer(0.5)

	// 让玩家移动到3,1点
	p.MoveTo(Vec2{3, 1})

	// 如果没有到达就一直循环
	for !p.IsArrived() {

		// 更新玩家位置
		p.Update()

		// 打印每次移动后的玩家位置
		fmt.Println(p.Pos())
	}

}
