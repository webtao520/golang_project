package main 

import (
	"fmt"
)

func main (){
	var arr []int 
	var num *int 
	fmt.Printf("%p\n",arr)
	fmt.Printf("%p",num)
}

/**
PS D:\goLang\github\golang_project\go容器\Go语言nil：空值\零值\Go语言nil：空值> go run 3.go
0x0
0x0
*/