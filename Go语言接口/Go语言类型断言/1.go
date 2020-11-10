package main 

import (
	"fmt"
)

func main(){
	 var x interface{}
	 x=10
	 value,ok:=x.(int)
	 fmt.Print(value,",",ok)
}

/**
PS D:\goLang\github\golang_project\Go语言接口\Go语言类型断言> go run 1.go
10,true
*/