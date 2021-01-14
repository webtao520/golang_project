package main 

import (
	"fmt"
	// json包实现了json对象的编解码，参见RFC 4627。Json对象和go类型的映射关系请参见Marshal和Unmarshal函数的文档。
	 "encoding/json"
)

// 结构体 与json
// 1. 序列化 把go 语言中的结构体变量 -- 》 json格式字符串
// 2. 反序列化 把json格式字符串 ---》 go 语言中能够识别的结构体变量

type person struct {
	Name string 	`json:"name" db:"name" ini:"name"`
	Age int `json:"age"`
}

func main (){
    p1:=person {
		 Name: "张三",
		 Age:9000,
	}
	// 序列化
	b,err:=json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
		return
	}
	fmt.Printf("%v\n",string(b)) // {"name":"张三","age":9000}

	// 反序列化
	str:=`{"name":"张三","age":9000}`
	var p2 person
	json.Unmarshal([]byte(str),&p2)
	fmt.Printf("%#v\n", p2) // main.person{Name:"张三", Age:9000}

	
}