package main

import "fmt"

// Go语言中推荐使用驼峰式命名
// var student_name string
var studentName string // Go语言中全局变量声明后必须使用，不使用就编译不过去

var demo int // Go语言中全局变量声明后必须使用，不使用就编译不过去

// var StudentName string

// 声明变量
// var name string
// var age int
// var isOk bool

// 批量声明
var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)

func main() {
	name = "理想"
	age = 16
	isOk = true
	// var heiheihei string
	// Go语言中非全局变量声明后必须使用，不使用就编译不过去
	fmt.Print(isOk) // 在终端中输出要打印的内容
	fmt.Println()
	fmt.Printf("name:%s\n", name) // %s:占位符 使用name这个变量的值去替换占位符
	fmt.Println(age)              // 打印完指定的内容之后会在后面加一个换行符
	// heiheihei = "嘿嘿嘿"

	// 声明变量同时赋值
	var s1 string = "whb"
	fmt.Println(s1)
	// 类型推导（根据值判断该变量是什么类型）
	var s2 = "20"
	fmt.Println(s2)
	// 简短变量声明，只能在函数里面用
	s3 := "哈哈哈"
	fmt.Println(s3)
	// s1 := "10" // 同一个作用域（{}）中不能重复声明同名的变量
	// 匿名变量是一个特殊的变量：_(后面学了函数再说)
}
