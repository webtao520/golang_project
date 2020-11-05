package main 


import (
	"fmt"
)

func main(){
	UserKey := fmt.Sprintln("eee", "seckill", 444)
	fmt.Println(UserKey)
}


/**
PS D:\goLang\github\golang_project\Go语言并发\高效地使用Go语言并发特性> go run 4.go
eee seckill 444


func Sprintln
func Sprintln(a ...interface{}) string

Sprintln采用默认格式将其参数格式化，串联所有输出生成并返回一个字符串。总是会在相邻参数的输出之间添加空格并在输出结束后添加换行符。
*/