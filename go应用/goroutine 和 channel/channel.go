package main

import "fmt"

func main() {
	// 演示管道使用
	// 1. 创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)
	// 2. 看看intChan 是什么
	fmt.Printf("intChan 的值=%v intChan 本身的地址=%p\n", intChan, &intChan) // 	intChan 的值=0xc000102080 intChan 本身的地址=0xc000006028

	// 3. 向管道中写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50
	//intChan<-98 // 注意点，当我们给管道写入数据时，不能超过其容量
	// 4. 看看管道的长度和cap(容量)
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan)) // channel len=3 cap=3

	// 5. 从管道中读取数据
	var num2 int
	var num3 int
	num2 = <-intChan
	num3 = <-intChan
	fmt.Println("将管道中值覆给num2，num2值为 : ", num2)                        // 将管道中值覆给num2，num2值为 :  10
	fmt.Println("将管道中值覆给num3，num3值为 : ", num3)                        // 将管道中值覆给num3，num3值为 :  211
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan)) // channel len=2 cap=3

	// 6. 在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock
	num4 := <-intChan
	num5 := <-intChan
	fmt.Println("num4=", num4, "num5=", num5) // fatal error: all goroutines are asleep - deadlock!

}

/*

PS D:\goLang\github\golang_project\go应用\goroutine 和 channel> go run .\channel.go
					intChan 的值=0xc000102080 intChan 本身的地址=0xc000006028
					channel len=3 cap=3
					将管道中值覆给num2，num2值为 :  10
					channel len=2 cap=3
*/
