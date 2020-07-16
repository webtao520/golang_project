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
	//cat11 := <-allChan
	//fmt.Println(cat11.Name) // .\allChanNew.go:24:19: cat11.Name undefined (type interface {} is interface with no methods)
	newCat := <-allChan
	//fmt.Printf("newCat=%T , newCat=%v\n", newCat, newCat) //newCat=main.Cat , newCat={tom 18}
	//fmt.Printf("newCat.Name=%v", newCat.Name) // .\allChanNew.go:27:37: newCat.Name undefined (type interface {} is interface with no methods)
	// 使用类型断言
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v", a.Name)

}
