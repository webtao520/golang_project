package main

import (
	"fmt"
	"strings"
)

// 字符串

func main (){
	//  \  本来时具有特殊含义的， 我应该告诉程序我写的 \ 就是一个单纯的 \ 
	path := "'D:\\Go\\src\\code.oldboyedu.com\\studygo\\day01'"
	fmt.Println(path) // 'D:\Go\src\code.oldboyedu.com\studygo\day01'

	s := "I'm ok"
	fmt.Println(s) // I'm ok

	// 多行字符串
	s2:= `
世情薄	 
			人情恶
	雨送黄昏花易落			
	`
	fmt.Println(s2)

	s3 := `D:\Go\src\code.oldboyedu.com\studygo\day01`  // 字符串原样输出
	fmt.Println(s3) // D:\Go\src\code.oldboyedu.com\studygo\day01

	// 字符串相关操作
	fmt.Println(len(s3))  // 42


	// 字符串拼接
	name:="理想"
	world := "大帅比"

	ss:=name + world
	fmt.Println(ss)

	/*
		func Sprintf(format string, a ...interface{}) string
		Sprintf根据 format 参数生成格式化的字符串并返回该字符串。
	*/
	ss1:=fmt.Sprintf("%s%s",name,world) // TODO
	fmt.Printf("%s%s",name,world)
	fmt.Println(ss1)

	// 字符串分割
	ret:=strings.Split(s3,"\\")
	fmt.Println(ret) // [D: Go src code.oldboyedu.com studygo day01]
	// 包含
	fmt.Println(strings.Contains(ss,"理性"))
	fmt.Println(strings.Contains(ss,"理想"))	
	// 前缀
	fmt.Println(strings.HasPrefix(ss,"理想"))
	fmt.Println(strings.HasSuffix(ss,"大"))

	s4:="abcdeb"
	/*
		func Index(s, sep string) int
		子串sep在字符串s中第一次出现的位置，不存在则返回-1。
	*/
	fmt.Println(strings.Index(s4,"c"))     // 2
	/*
		func LastIndex(s, sep string) int
		子串sep在字符串s中最后一次出现的位置，不存在则返回-1。
	*/
	fmt.Println(strings.LastIndex(s4,"b")) // 5

	fmt.Println(strings.LastIndex(s4,"gg")) // -1
}
