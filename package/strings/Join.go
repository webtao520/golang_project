package main

import (
	"fmt"
	"strings"
)

func main (){
	//将一系列字符串连接为一个字符串，之间用sep来分隔。
	s:=[]string{"foo","bar","baz"}
	fmt.Println(strings.Join(s,""))
}

/*
PS D:\goLang\github\golang_project\package\strings> go run Join.go
foo, bar, baz
*/