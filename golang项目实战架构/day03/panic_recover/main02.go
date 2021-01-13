package main 


import (
	"fmt"
)

func funcA() {
	fmt.Println("func A")
}

//可以通过recover将程序恢复回来，继续往后执行。
func funcB(){
    defer func(){
		err:=recover()
		// 如果程序出现了panic 错误，可以通过recover恢复过来, recover 收集错误
		if err != nil{
			fmt.Println(err)
			fmt.Println("recover in B...")	
		}
	}()
	panic("panic in B....")
}

func funcC() {
	fmt.Println("func C...")
}


func main(){
	funcA()
	funcB()
	funcC()
}



/**
注意：

recover()必须搭配defer使用。
defer一定要在可能引发panic的语句之前定义。

PS D:\goLang\github\golang_project\golang项目实战架构\day03\panic_recover> go run .\main02.go

func A
panic in B....
recover in B...
func C...
*/