package main 

import (
	"fmt"
)

func main (){
	a:=[2]int{1,2}
	b:=[...]int{1,2}
	c:=[2]int{1,3}
	fmt.Println(a==b,a==c,b==c)
}

/**
PS D:\golang\github\golang_project\go容器\Go语言数组> go run 5.go
true false false
*/