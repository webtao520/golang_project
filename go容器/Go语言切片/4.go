package main

import (
	"fmt"
)

func main (){
	a:=[]int{1,2,3}
	fmt.Println(a[0:0])
}

/**
PS D:\golang\github\golang_project\go容器\Go语言切片> go run 4.go
[]
*/