package main 

import (
	"fmt"
	"container/list"
)

func main(){
	l:=list.New()
	l.PushBack("fist")
	l.PushFront(67)
	fmt.Println(l)
}


/*
PS D:\goLang\github\golang_project\go容器\Go语言list（列表）> go run 1.go
&{{0xc00006c390 0xc00006c360 <nil> <nil>} 2}
*/