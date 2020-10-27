package main 

import (
	"fmt"
)

func main (){
	var sum *int
	fmt.Println(sum) //零值是Go语言中变量在声明之后但是未初始化被赋予的该类型的一个默认值
	sum = new(int)    // 分配空间
	fmt.Println(sum) // 获取变量地址
	//*sum = 98
	//fmt.Println(*sum)
}


/**
PS D:\goLang\github\golang_project\go容器\Go语言make和new关键字的区别及实现原理> go run  1.go
0xc00000a0b0
98
*/