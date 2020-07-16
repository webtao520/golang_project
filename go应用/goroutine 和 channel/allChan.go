//  创建一个allChan ，最多可以存放10个任意数据类型变量，演示写入和读取的用法
package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age  int
}

func main() {
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)
	cat1 := Cat{Name: "tom", Age: 18}
	cat2 := Cat{Name: "tom~", Age: 180}
	allChan <- cat1
	allChan <- cat2
	allChan <- 10
	allChan <- "jack"
	// 取出
	cat11 := <-allChan
	cat22 := <-allChan
	v1 := <-allChan
	v2 := <-allChan
	fmt.Println(cat11, cat22, v1, v2)
}

/*
PS D:\goLang\github\golang_project\go应用\goroutine 和 channel> go run .\allChan.go
		{tom 18} {tom~ 180} 10 jack
*/
