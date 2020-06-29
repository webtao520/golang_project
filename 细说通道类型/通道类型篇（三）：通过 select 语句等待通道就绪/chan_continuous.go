/*
另外，需要注意的是，select 语句只能对其中的每一个 case 表达式各求值一次，如果我们想连续操作其中的通道的话，
需要通过在 for 语句中嵌入 select 语句的方式来实现：
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	chs := [3]chan int{
		make(chan int, 3),
		make(chan int, 3),
		make(chan int, 3),
	}

	index1 := rand.Intn(3) // 随机生成0-2之间的数字
	fmt.Printf("随机索引/数值: %d\n", index1)
	chs[index1] <- rand.Int() // 向通道发送随机数字

	index2 := rand.Intn(3)
	fmt.Printf("随机索引/数值: %d\n", index2)
	chs[index2] <- rand.Int()

	index3 := rand.Intn(3)
	fmt.Printf("随机索引/数值: %d\n", index3)
	chs[index3] <- rand.Int()

	// 哪一个通道中有值，哪个对应的分支就会被执行
	for i := 0; i < 3; i++ {
		select {
		case num, ok := <-chs[0]:
			if !ok {
				break
			}
			fmt.Println("第一个条件分支被选中: chs[0]=>", num)
		case num, ok := <-chs[1]:
			if !ok {
				break
			}
			fmt.Println("第二个条件分支被选中: chs[1]=>", num)
		case num, ok := <-chs[2]:
			if !ok {
				break
			}
			fmt.Println("第三个条件分支被选中: chs[2]=>", num)
		default:
			fmt.Println("没有分支被选中")
		}
	}
}

/*
但这时要注意，简单地在 select 语句的分支中使用 break 语句，只能结束当前的 select 语句的执行，
而并不会对外层的 for 语句产生作用，如果 for 循环本身没有退出机制的话会无休止地运行下去。
*/
