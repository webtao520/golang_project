/*
	1. 统计字符串的长度，按字节len(str)
	2. golang的编码统一为utf-8 (ASCII字符（字母和数字）占一个字节，汉字占用3个字节)
*/
package main 

import (
	"fmt"
)


func main (){
  str:="hello北"
  fmt.Println(len(str))	
}