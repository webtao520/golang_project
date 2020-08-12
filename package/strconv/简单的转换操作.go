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

/**
但注意:
		1. 不是所有数据类型都能转换，例如字母格式的string类型 "abcd" 转换为int 肯定会失败
		2. 低精度转换为高精度时是安全的，高精度的值转换为低精度时会丢失精度，例如int32转换为int16。 float32转换为int
		3. 这种简单的转换方式不能对int(float)string 进行互转，要跨大类型转换，可以使用strconv包提供的函数
*/
