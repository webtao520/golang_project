package main

import (
	"fmt"
)

func main() {
	/*
			str:="abc"
		    s2:=[]byte(str)
		    s2[0]='b'
			fmt.Println(string(s2)) //bbc
	*/
	str := "白猫"
	s2 := []rune(str)
	s2[0] = '黑'
	fmt.Println(string(s2))
}

// 上面使用到的 string()，表示强制类型转换，转换为字符串。
