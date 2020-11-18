package main 


import (
	"fmt"
)

func printType (v interface{}){
	switch v.(type) {
	case int:
		fmt.Println(v, "is int")
	case string:
		fmt.Println(v, "is string")
	case bool:
		fmt.Println(v, "is bool")
	}
}

func main (){
	printType(1024)
	printType("pig")
	printType(true)
}


/**
PS D:\goLang\github\golang_project\Go语言接口\Go语言类型分支（switch判断空接口中变量的类型）> go run 1.go
1024 is int
pig is string
true is bool
*/