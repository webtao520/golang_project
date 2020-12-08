package main

import (
	"encoding/json"
	"fmt"
)

// 声明技能结构体
type Skill struct {
	Name string `json:"name"`
	Level int  `json:"level"`
}

// 声明角色结构体
type Actor  struct {
	Name string
	Age int 
	Skills []Skill
}

func main (){
	// 填充基本角色数据
	a:=Actor{
		 Name:"cow boy",
		 Age:37,
		 Skills: []Skill{
			{Name: "Roll and roll", Level: 1},
            {Name: "Flash your dog eye", Level: 2},
            {Name: "Time to have Lunch", Level: 3},
		 },
	}
	// 转化为json
	result,err:=json.Marshal(a)
	if err !=nil {
		fmt.Println(err)
	}
	// 将字节型（字符型） byte []  uint8 转化为字符串  8 位无符号整型
	jsonStringData :=string(result)
	fmt.Println(jsonStringData)

}


/**
PS D:\goLang\github\golang_project\Go语言结构体\示例：将结构体数据保存为JSON格式数据> go run 2.go
{"Name":"cow boy","Age":37,"Skills":[{"name":"Roll and roll","level":1},{"name":"Flash your dog eye","level":2},{"name":"Time to have Lunch","level":3}]}
*/