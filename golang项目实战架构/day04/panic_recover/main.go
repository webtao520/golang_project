package main 

import (
	"fmt"
)

// panic 和 recover
func f1() {
	defer func (){
		err:=recover() // // 收集当前的错误
		fmt.Println("gogogo....")
		fmt.Println(err)	
	}()
	panic("犯了不可原谅的错误")
	fmt.Println("f1")
}

func f2(){
	fmt.Println("f2")
}

func main(){
	f1()
	f2()
}

/**
PS D:\goLang\github\golang_project\golang项目实战架构\day04\panic_recover> go run main.go
gogogo....
犯了不可原谅的错误
f2
*/

