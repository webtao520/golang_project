package main

import (
	"fmt"
	"sort"
)

func main (){
	scene :=make(map[string]int)
	//准备map数据
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	// 声明一个切片保存map数据
	var sceneList  []string
	//将map数据遍历复制切片中
	for k :=range scene {
		sceneList=append(sceneList,k)
	}

	// 对切片进行排序
	sort.Strings(sceneList)  // sort.Strings 的作用是对传入的字符串切片进行字符串字符的升序排列

	//输出
	fmt.Println(sceneList)
}

/**
PS D:\goLang\github\golang_project\go容器\Go语言遍历map> go run 2.go
[brazil china route]
*/