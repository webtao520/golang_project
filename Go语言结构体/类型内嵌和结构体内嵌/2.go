package main 

import (
	"fmt"
)

type A struct {
	 ax,ay int 
}

type B struct {
	A
	bx,by float32
}

func main (){
	b:=B{A{1,2},3.0,4.0}
	fmt.Println(b.ax,b.ay,b.bx,b.by)
	fmt.Println(b.A)
}

/**
PS D:\goLang\github\golang_project\Go语言结构体\类型内嵌和结构体内嵌> go run 2.go
1 2 3 4
{1 2}
*/