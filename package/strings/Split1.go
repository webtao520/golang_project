package main

import (
	"fmt"
	"strings"
)

/**
	func Split
	func Split(s, sep string) []string

用去掉s中出现的sep的方式进行分割，会分割到结尾，
并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，即使两个sep相邻，
也会进行两次切割）。
如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。	
*/

func main (){
	// %q      单引号围绕的字符字面值，由Go语法安全地转义 Printf("%q", 0x4E2D)  '中'
	fmt.Printf("%q\n",strings.Split("a,b,c",",")) // ["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) // ["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split(" xyz ", "")) // [" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins")) [""]
}