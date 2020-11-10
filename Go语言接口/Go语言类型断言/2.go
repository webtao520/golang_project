package main 

import (
	"fmt"
)

func main(){
	var x interface{}
	x="hello"
	value:=x.(int)
	fmt.Println(value)
}

/**
PS D:\goLang\github\golang_project\Go语言接口\Go语言类型断言> go run 2.go
panic: interface conversion: interface {} is string, not int
*/