package main 

import  (
	"fmt"
)

// panic 和   recover
func funcA(){
	fmt.Println("a")
}

/**
	Go语言中目前（Go1.12）是没有异常机制，但是使用panic/recover模式来处理错误。
	panic可以在任何地方引发，但recover只有在defer调用的函数中有效。
*/

func funcB(){
	// 刚刚打开数据库连接
	defer func(){
		 err := recover()
		 fmt.Println(err)
		 fmt.Println("释放数据库连接...")
	}()
	panic("出现了严重的错误 !!!")
	fmt.Println("b")
}

func funcC(){
	fmt.Println("c")
}

func main (){
	funcA()
	funcB()
	funcC()
}

/**
PS D:\goLang\github\golang_project\golang项目实战架构\day03\panic_recover> go run main.go
a
出现了严重的错误 !!!
释放数据库连接...
c
*/