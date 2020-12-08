package main

import (
	"encoding/json"
	"fmt"
)

// 声明技能结构体
type Skill struct {
	Name  string
	Level int
}

// 声明角色结构体
type Actor struct {
	Name   string
	Age    int
	Skills []Skill
}

func main() {
	// 填充基本角色数据
	a := Actor{
		Name: "cow boy",
		Age:  37,
		Skills: []Skill{
			{Name: "Roll and roll", Level: 1},
			{Name: "Flash your dog eye", Level: 2},
			{Name: "Time to have Lunch", Level: 3},
		},
	}
	// 转化为json
	result, err := json.Marshal(a) //
	fmt.Printf("%T", result)       // []uint8 字节型(字符类型)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	jsonStringData := string(result)
	// {"Name":"cow boy","Age":37,"Skills":[{"Name":"Roll and roll","Level":1},{"Name":"Flash your dog eye","Level":2},{"Name":"Time to have Lunch","Level":3}]}
	fmt.Println(jsonStringData)
}
