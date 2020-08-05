/**

sice和数组很像，但是slice可以添加元素

	slice中文可以称为“切片”。是Go语言为处理同类型数据序列提供的一个高效且方便的方式。是在数组上抽象的一个数据类型。
	切片的使用相当的广泛。
	切片可以直接赋值，不用提前指定元素的个数。切片还可以动态的增加元素。
	一个slice也就是切片的声明是这样的

	s := []int{1,1,2,3,5,8,13,21,34}
	也可以使用make创建，make创建slice时需要指定元素个数。

	a := make([]int,0)
	先指定为0个

*/

/*
func main() {
	s := []int{1, 1, 2, 3, 5, 8, 13, 21, 34}
	a := make([]int, 10)
	fmt.Println(s)
	fmt.Println(a)
}
*/

/**
	运行结果如下
	PS D:\goLang\github\golang_project\go应用\slice> go run .\slice和数组很像，但是slice可以添加元素.go
				[1 1 2 3 5 8 13 21 34]
				[0 0 0 0 0 0 0 0 0 0]



slice添加元素是这样写的  a = append(a, 55, 89)


这个append给slice a添加了2个整数。但是，其实 = 左边和右边的a并不是同一个a了。
为了验证这两个a已经发生了根本的变化，我们来取得append前后a的内存地址做比较。
不过，因为我们之前声明a的时候，没有给a元素个数，那么它那个时候还没有分配内存。我们先修改一下

*/

/*
func main() {
	//a := make([]int, 0)
	//fmt.Println(a)
	// 给a元素个数, 分配内存
	//a := make([]int, 1)
	//fmt.Println(a)

	s := []int{1, 1, 2, 3, 5, 8, 13, 21, 34}
	a := make([]int, 1)
	fmt.Println(s)
	fmt.Println(a)

	fmt.Println(&a[0])
	a = append(a, 55, 89)
	fmt.Println(a)
	fmt.Println(&a[0])

}
*/

/*
运行的结果是这样的

		PS D:\goLang\github\golang_project\go应用\slice> go run .\slice和数组很像，但是slice可以添加元素.go
		[1 1 2 3 5 8 13 21 34]
		[0]
		0xc00000a0b0
		[0 55 89]
		0xc000010400

		注意输出的运行结果中[0 55 89]前后两个地址的不同。当然，在你的机器上的地址很可能与我的不同，
		那是因为每次运行时，环境不同就可能有不同的内存分配过来。
		既然两个a是不同的变量（因为a的内存地址会发生变化，通过地址访问会掉入逻辑错误的坑里），
		那么，下面的代码就容易理解了。
*/

package main

import "fmt"

func main() {
	a := make([]int, 1)
	b := make([]int, 0)
	b = append(a, 144)
	fmt.Println(&b[0]) // 0xc00000a0c0

	// 然后你再给 b = append()增加一个元素看看效果。然后再减掉刚增加的这个看看有什么变化
	b = append(a, 144, 233)
	fmt.Println(&b[0]) // 0xc0000103e0
}

/*
	是的，看到了b的内存地址的变化了。这是因为原来的slice对应的容量不够了，
	随着slice元素的增加，又重新分配了内存地址。
	所以，这里有一个结论，slice在内存中保存的位置并不是稳定，没什么必要的话，
	不要用内存地址或者指针玩slice，小心有坑。
*/
