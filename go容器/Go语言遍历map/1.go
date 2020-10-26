package main 

import (
	"fmt"
)

func main(){
	scene :=make(map[string]int)
	scene["route"]=66
	scene["brazil"]=4
	scene["china"]=960
	for k,v:=range scene{
		fmt.Println(k,v)
	}
}

/**
	PS D:\goLang\github\golang_project\go容器\Go语言遍历map> go run 1.go
	route 66
	brazil 4
	china 960
*/