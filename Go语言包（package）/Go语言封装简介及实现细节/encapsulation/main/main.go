package main 


import (
	"fmt"
	"../model"
)

func main (){
	p:=model.NewPerson("smith")
    p.SetAge(18)
    p.SetSal(5000)
    fmt.Println(p)
    fmt.Println(p.Name, " age =", p.GetAge(), " sal = ", p.GetSal())
}

/**
	PS D:\goLang\github\golang_project\Go语言包（package）\Go语言封装简介及实现细节\encapsulation\main> go run .\main.go
	&{smith 18 5000}
	smith  age = 18  sal =  5000
*/