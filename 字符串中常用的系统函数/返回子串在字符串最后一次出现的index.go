package main 

import (
	"fmt"
	"strings"
)

func main (){
	index:=strings.Index("LAMP_go_lang","go")
	// 
	fmt.Printf("index=%v\n",index)
}

/*
PS D:\goLang\github\golang_project\字符串中常用的系统函数> go run .\返回子串在字符串最后一次出现的index.go
														index=5
*/