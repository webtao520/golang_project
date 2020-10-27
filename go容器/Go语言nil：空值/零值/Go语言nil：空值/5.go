package main 

import (
	"fmt"
)

func main(){
	var s1 []int 
	var s2 []int 
	fmt.Printf(s1 == s2)
}

/*
PS D:\goLang\github\golang_project\go容器\Go语言nil：空值\零值\Go语言nil：空值> go run 5.go
# command-line-arguments
.\5.go:10:16: invalid operation: s1 == s2 (slice can only be compared to nil)
*/