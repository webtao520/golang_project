/*
		package builtin
		import "builtin"
		builtin 包为Go的预声明标识符提供了文档。此处列出的条目其实并不在builtin 包中，
		对它们的描述只是为了让 godoc 给该语言的特殊标识符提供文档。

		1. len: 用来求长度，比如 strings,array,slice,map,channel
		2. new: 用来分配内存，主要用来分配值的类型，比如int ,float32,struct。。。返回的是指针
		3. make: 用来分配内存主要用来分配引用类型，比如channel,map,slice
		下面是举例说明new的使用

		相应的文档 https://studygolang.com/pkgdoc   ----> builtin
*/

package main 

import (
	"fmt"
)

func main () {
   num1:=100
   fmt.Printf("num1的类型%T, num1的值=%v, num1的地址%v\n",num1,num1,&num1)

   num2:=new(int)
   *num2=100
   fmt.Printf("num2的类型%T, num2的值=%v, num2的地址%v\n num2这个指针，指向的值=%v",num2,num2,&num2,*num2)
}

/*
	PS D:\goLang\github\golang_project\go应用\内置函数> go run .\len.go
				num1的类型int, num1的值=100, num1的地址0xc00000a0b0
				num2的类型*int, num2的值=0xc0000a00b0, num2的地址0xc0000ca020 num2这个指针，指向的值=100
*/