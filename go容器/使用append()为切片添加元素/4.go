package main

import (
	"fmt"
)

func main() {
	var a []int
	a = append(a[:i], append([]int{x}, a[i:]...)...)       // 在第i个位置插入x
	a = append(a[:i], append([]int{1, 2, 3}, a[i:]...)...) // 在第i个位置插入切片
	fmt.Println(a)
}

//每个添加操作中的第二个 append 调用都会创建一个临时切片，并将 a[i:] 的内容复制到新创建的切片中，然后将临时创建的切片再追加到 a[:i] 中。
