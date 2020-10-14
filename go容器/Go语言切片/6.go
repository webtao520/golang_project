package main 

import (
	"fmt"
)

func main (){
	a:=make([]int,2)
	b:=make([]int,2,10)
	fmt.Println(a,b)
	fmt.Println(len(a),len(b))
}

/**
PS D:\golang\github\golang_project\go容器\Go语言切片> go run 6.go
[0 0] [0 0]
2 2
*/