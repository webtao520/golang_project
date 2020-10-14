package main 

import (
	"fmt"
)

func main () {
	var a = [3]int{1,2,3}
	fmt.Println(a,a[1:2])
}

/**
  PS D:\golang\github\golang_project\go容器\Go语言切片> go run 1.go
	[1 2 3] [2]
*/