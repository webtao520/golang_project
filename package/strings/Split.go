package main 

import (
	"fmt"
	"strings"
)

func main () {
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
}

/*
		func Split
func Split(s, sep string) []string
用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。

		PS D:\goLang\github\golang_project\package\strings> go run .\Split.go
		["a" "b" "c"]
		["" "man " "plan " "canal panama"]
		[" " "x" "y" "z" " "]
		[""]

*/