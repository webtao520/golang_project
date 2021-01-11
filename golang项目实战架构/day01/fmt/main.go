package main 

import (
	"fmt"
)

// fmt 占位符
func main(){
	var n = 100

	fmt.Printf("%T\n",n)  	// 查看类型
	fmt.Printf("%v\n",n) // 相应值的默认格式。
	fmt.Printf("%b\n",n) //   二进制表示 
	fmt.Printf("%d\n", n) // 十进制表示  
	fmt.Printf("%o\n", n) //  八进制表示  
	fmt.Printf("%x\n", n) //  十六进制表示  

	var s = "Hello 沙河！"
	fmt.Printf("字符串: %s\n", s)  // 相应值的默认格式。
	fmt.Printf("字符串: %v\n", s) 
	fmt.Printf("字符串：%#v\n", s) // 相应值的go语法表示  (   字符串："Hello 沙河！"   )

}

/**
		int
		100
		1100100
		100
		144
		64
		字符串: Hello 沙河！
		字符串: Hello 沙河！
		字符串："Hello 沙河！"
*/