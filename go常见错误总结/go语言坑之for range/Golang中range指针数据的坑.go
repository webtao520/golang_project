/*
在Golang中使用for range语句进行迭代非常的便捷，但在涉及到指针时就得小心一点了。

下面的代码中定义了一个元素类型为*int的通道ch：
*/

package main

import "fmt"

func main() {
	ch := make(chan *int, 5)
	input := []int{1, 2, 3, 4, 5}
	go func() {
		for k, _ := range input {
			ch <- &input[k]
		}
		close(ch)
	}()

	for v := range ch {
		// fmt.Println(*v)
		fmt.Println(*v)
	}
}

/**
PS D:\goLang\github\golang_project\go常见错误总结\go语言坑之for range> go run .\Golang中range指针数据的坑.go
			5
			5
			5
			5
			5
很明显，程序并没有达到预期的结果，那么问题出在哪里呢？我们将代码稍作修改：

PS D:\goLang\github\golang_project\go常见错误总结\go语言坑之for range> go run .\Golang中range指针数据的坑.go
0xc00000a0b0
0xc00000a0b0
0xc00000a0b0
0xc00000a0b0
0xc00000a0b0

可以看到，5次输出变量v（*int）都指向了同一个地址，返回去检查一下发送部分代码：

for _, v := range input {
    ch <- &v
}

问题正是出在这里，在for range语句中，v变量用于保存迭代input数组所得的值，但是v只被声明了一次，此后都是将迭代input出的值赋值给v，
v变量的内存地址始终未变，这样再将v的地址发送给ch通道，发送的都是同一个地址，当然无法达到预期效果。

解决方案是，引入一个中间变量，每次迭代都重新声明一个变量temp，赋值后再将其地址发送给ch：

for _, v := range input {
    temp := v
    ch <- &temp
}

抑或直接引用数据的内存（推荐，无需开辟新的内存空间）

for k, _ := range input {
    c <- &input[k]
}

再次运行，就可看到预期的效果。以上方案是用于讨论range语句带来的问题，当然，平时还是尽量避免使用指针类型的通道。

*/
