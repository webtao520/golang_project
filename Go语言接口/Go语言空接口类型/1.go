package main 

import (
	"fmt"
)

func main (){
	// 声明a变量。类型int，初始值为1 
	var a int =1 
	// 声明i变量，类型为interface{},初始值为a，此时i的值变为1 
	var i interface {} = a  
	// 将 a 的值赋值给 i 时，虽然 i 在赋值完成后的内部值为 int，但 i 还是一个 interface{} 类型的变量。
	// 类似于无论集装箱装的是茶叶还是烟草，集装箱依然是金属做的，不会因为所装物的类型改变而改变。
	//fmt.Println(i)

	// 声明b变量，尝试赋值i
	//var b int =i  // 要进行类型断言

	var b int = i.(int)
	fmt.Println(b)
}