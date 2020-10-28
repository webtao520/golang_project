package main

import "fmt"

// 实际打印的函数
func rawPrint(rawList ...interface{}) {

	// 遍历可变参数切片
	for _, a := range rawList {

		// 打印参数
		fmt.Println(a)
	}
}

// 打印函数封装
func print(slist ...interface{}) {

	// 将slist可变参数切片完整传递给下一个函数
	rawPrint(slist...)
}

func main() {

	print(1, 2, 3)
}
