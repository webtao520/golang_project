package main 


import (
	"fmt"
)

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}

// 程序运行期间funcB中引发了panic导致程序崩溃，异常退出了
func main() {
	funcA()
	funcB()
	funcC()
}

/**
PS D:\goLang\github\golang_project\golang项目实战架构\day03\panic_recover> go run .\main01.go
func A
panic: panic in B

goroutine 1 [running]:
main.funcB(...)
        D:/goLang/github/golang_project/golang项目实战架构/day03/panic_recover/main01.go:13
main.main()
        D:/goLang/github/golang_project/golang项目实战架构/day03/panic_recover/main01.go:23 +0x9d
exit status 2
*/