package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%T", nil)
	print(nil)
}

/*
	PS D:\goLang\github\golang_project\go容器\Go语言nil：空值\零值\Go语言nil：空值> go run 2.go
	# command-line-arguments
	.\2.go:9:7: use of untyped nil
*/
