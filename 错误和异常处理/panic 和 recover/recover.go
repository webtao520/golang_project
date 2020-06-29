/*
# recover #
		我们还可以通过 recover() 函数对 panic 进行捕获和处理，从而避免程序崩溃然后直接退出，
		实现类似 PHP 中 try...catch...finally 的功能，由于执行到抛出 panic 的问题代码时，
		会中断后续其他代码的执行，所以，显然这个 panic 的捕获和其他代码的恢复执行需要放到 defer 语句中完成。
		下面我们引入 recover() 函数来改写上述代码如下：
*/

package main

import (
    "fmt"
)

func divide() {

    defer func() {
        if err := recover(); err != nil {
            fmt.Printf("Runtime panic caught: %v\n", err)  // %v	按值的本来值输出
        }
	}()
	

    var i = 1
    var j = 0
    k := i / j
    fmt.Printf("%d / %d = %d\n", i, j, k)
}

func main() {
    divide()
    fmt.Println("divide方法调用完毕，回到main函数")
}