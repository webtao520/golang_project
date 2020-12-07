package main

import (
	"fmt"
)

// 车轮
type Wheel struct {
	Size int
}

// 车
type Car struct {
	Wheel
	//引擎 匿名结构体不使用type 声明
	Engine struct {
		Power int // 功率
		Type  string
	}
}

func main() {
	c := Car{
		// 初始化轮子
		Wheel: Wheel{
			Size: 18,
		},
		// 初始化引擎
		Engine: struct {
			Power int
			Type  string
		}{
			Type:  "1.4T",
			Power: 143,
		},
	}
	fmt.Printf("%+v\n", c)
}

/**
PS D:\goLang\github\golang_project\Go语言结构体\初始化内嵌结构体> go run 2.go
{Wheel:{Size:18} Engine:{Power:143 Type:1.4T}}
*/
