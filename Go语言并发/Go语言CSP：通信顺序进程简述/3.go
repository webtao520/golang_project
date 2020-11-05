package main 

import (
	"fmt"
)

func main(){
	done:=make(chan int)
	
	go func (){
		fmt.Println("golang")
		data:=<-done
		fmt.Println(data)
	}()
	
	 done<-1
}

/**
	PS D:\goLang\github\golang_project\Go语言并发\Go语言CSP：通信顺序进程简述> go run 3.go
	golang
	1
*/