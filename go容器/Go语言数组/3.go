package main 

import (
	"fmt"
)

func  main (){
	q:=[...]int{1,2,3}
	fmt.Printf("%T\n",q)
}

/**
	PS D:\golang\github\golang_project\go容器\Go语言数组> go run 3.go
	[3]int
*/