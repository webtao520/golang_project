package main 

import (
	"fmt"
)

func main (){
	q:=[3]int{1,2,3}
	q=[4]int{1,2,3,4}
	fmt.Println(q)
}

/**
PS D:\golang\github\golang_project\go容器\Go语言数组> go run 4.go
# command-line-arguments
.\4.go:9:3: cannot use [4]int literal (type [4]int) as type [3]int in assignment
*/