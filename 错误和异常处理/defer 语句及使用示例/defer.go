/*
	defer func (){
		// 执行复杂的清理工作 ...
	}

	另外，一个函数/方法中可以存在多个 defer 语句，defer 语句的调用顺序遵循先进后出的原则，即最后一个 defer 语句将最先被执行，
	相当于「栈」结构，如果在循环语句中包含了 defer 语句，则对应的 defer 语句执行顺序依然符合先进后出的规则。
	由于 defer 语句的执行时机和调用顺序，所以我们要尽量在函数/方法的前面定义它们，以免在后面执行时漏掉，尤其是运行时抛出错误会中断后面代码的执行，
	也就感知不到后面的 defer 语句。
	下面我们看一个简单的示例 defer.go：
*/

package main 


import (
	"fmt"
)

func printError(){
	fmt.Println("娄底执行")
}

func main (){

	 // 栈结构，先进后出
	defer printError()
	defer func (){
		fmt.Println("除数不能为0！")
	}()
	
	var i =1 
	var j=0
	var k=i/j
	fmt.Printf("%d / %d = %d\n",i,j,k)
}

/*
			这段代码中，我们定义了两个 defer 语句，并且是在函数最顶部，以确保异常情况下也能执行。

			在函数正常执行的情况下，这两个 defer 语句会在最后一条打印语句执行完成后先执行第二条 defer 语句，再执行第一条 defer 语句

			1 / 1 = 1
			除数不能为0！
			娄底执行

			而如果我们把 j 的值设置为 0，则函数会抛出 panic：
			====================================================
			除数不能为0！
			娄底执行
			panic: runtime error: integer divide by zero

			goroutine 1 [running]:
			main.main()
					D:/goLang/github/golang_project/错误和异常处理/defer 语句及使用示例/defer.go:34 +0x58
			exit status 2
			===================================================

			表示除数不能为零，这个时候，也会执行 defer 语句，底层的逻辑是在执行 var k = i / j 这条语句时，遇到除数为0，
			则抛出异常 panic，然后立即中断当前函数 main 的执行（后续语句都不再执行），
			并按照先进后出顺序依次执行定义在当前函数中的 defer 语句，最后打印出 panic 日志及错误信息。
			关于 panic 及其内部执行逻辑，学院君将在下一篇教程给大家介绍。

			总结一下，Go 语言的 defer 语句相当于 PHP 中析构函数和 finally 语句的功效，
			常用于定义兜底逻辑，在函数执行完毕后或者运行抛出 panic 时执行，如果一个函数定义了多个 defer 语句，则按照先进后出的顺序执行。
*/