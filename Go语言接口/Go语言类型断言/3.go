package main 

import (
	"fmt"
)

func main(){
	var a int 
	a = 10
	getType(a)
}


func getType(a interface{}) {
    switch a.(type) {
    case int:
        fmt.Println("the type of a is int")
    case string:
        fmt.Println("the type of a is string")
    case float64:
        fmt.Println("the type of a is float")
    default:
        fmt.Println("unknown type")
    }
}

/**
PS D:\goLang\github\golang_project\Go语言接口\Go语言类型断言> go run 3.go
the type of a is int
*/