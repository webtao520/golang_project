/*
	空接口 (interface{}) 不包含任何的方法，正因为如此，所有的类型都实现了空接口，因此空接口可以存储任意类型的数值，

	当函数可以接受任意的对象实例时，我们会将其声明为interface{} 最定性的例子是标准库fmt中PrintXXX系列函数例如
	 func Printf(fmt string, args ...interface{})
	 func Println(args ...interface{})
*/

package main

import "fmt"

func main() {
	var v1 interface{} = 1     // 将int类型赋值给interface{}
	var v2 interface{} = "abc" // 将string类型赋值给interface{}
	var v3 interface{} = &v2   // 将*interface{}类型赋值给interface{}
	var v4 interface{} = struct{ X int }{1}
	var v5 interface{} = &struct{ X int }{1}

	fmt.Println(v1, v2, v3, v4, v5)

	//fmt.Println(args ...interface{ })
}

/*
PS D:\goLang\github\golang_project\go面向对象特性\接口> go run .\interface05_空接口.go
					1 abc 0xc0001021e0 {1} &{1}
*/
