/*
	Go不会对数据进行隐式的类型转换，只能手动去执行转换操作。
	简单的转换操作
	转换数据类型的方式很简单。

	valueOfTypeB = typeB(valueOfTypeA)
*/

package main

import "fmt"

func main() {
	/*
		// 浮点数
		a := 5.0
		// 转换为int类型
		b := int(a)
		fmt.Printf("%T", b)
	*/

	// IT类型的底层是int类型
	type IT int

	// a的类型为IT，底层是int
	var a IT = 5

	// 将a(IT)转换为int，b现在是int类型
	b := int(5)

	// 将b(int)转换为IT，c现在是IT类型
	c := IT(b)

	fmt.Printf("%T\n", c)
	fmt.Println(a, b)

}

/*
##################### 运行结果 #######################

PS D:\goLang\github\golang_project\package\strconv> go run 简单的转换操作.go
					int

PS D:\goLang\github\golang_project\package\strconv> go  run 简单的转换操作.go
					main.IT
					5 5
*/
