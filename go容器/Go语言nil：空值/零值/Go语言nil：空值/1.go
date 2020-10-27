package main 


import (
	"fmt"
)

func main(){
	fmt.Println(nil==nil)
}



/*
	PS D:\goLang\github\golang_project\go容器\Go语言nil：空值\零值\Go语言nil：空值> go run 1.go
	# command-line-arguments
	.\1.go:9:17: invalid operation: nil == nil (operator == not defined on nil)
*/