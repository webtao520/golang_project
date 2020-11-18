package main 


import (
	"fmt"
)

func main (){
	// a 保存整型
	var a interface {} =100
	// b 保存字符串
	var b interface {} ="hi"
	// 两个空接口不想等
	fmt.Println(a == b)
}