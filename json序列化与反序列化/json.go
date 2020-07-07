package main 

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string
	Age int 
	Birthday string
	Sal float64
	SKill string
}


func testStruct(){
	monster:=Monster{
	   Name:"小狗",
	   Age:500,
	   Birthday: "2020-11-11",
	   Sal:8000.0,
	   SKill: "咬人",	
	}
	// 将monster 序列化
	data,err:=json.Marshal(&monster)
	if err!=nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后的结果
	fmt.Printf("monster 序列化后=%v\n",string(data))
}


// 将map 进行序列化
func  testMap (){
	var a map[string]interface{}
	a=make(map[string]interface{})
	a["name"]="红孩儿"
	a["age"]=45
	a["addresss"]="红彦东"
	data,err:=json.Marshal(a)
	if err!=nil {
		fmt.Printf("序列化错误 err=%v\n",err)
	}
	fmt.Printf("a map 序列化后=%v\n",string(data))
}

func testSlice(){
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1=make(map[string]interface{})
	m1["name"]="jack"
	m1["age"]="7"
	m1["address"]="上海"
	slice=append(slice,m1)

	var m2 map[string]interface{}
	m2=make(map[string]interface{})
	m2["name"]="tom"
	m2["age"]="20"
	m2["address"]=[2]string{"墨西哥","日本"}
	slice=append(slice,m2)

	data,err:=json.Marshal(slice)
	if err !=nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
   fmt.Printf("slice 学历恶化后=%v\n",string(data))
}

// 对基本数据类型进行序列化 ,基本类型进行序列化意义不大
func testFloat64(){
	var num1  float64 =2345.34
	
	data,err:=json.Marshal(num1)
	if err !=nil {
		fmt.Printf("序列化错误 err=%v\n",err)
	}
	fmt.Printf("num1 序列化后=%v\n",string(data))
}


// json 序列化时指，将有key-value结构的数据类型(比如结构体，map,切片)序列化成json字符串的操作
func main (){
	// 演示将结构体 map 切片进行序列化
   //testStruct()
   //testMap()
   //testSlice()
   testFloat64()
}