//sort.Sort是递增排序，如果要实现递减排序，用sort.Reverse
package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{4, 3, 2, 1, 5, 9, 8, 7, 6}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	/*
		IntSlice给[]int添加方法以满足Interface接口，以便排序为递增序列。
		Reverse包装一个Interface接口并返回一个新的Interface接口，对该接口排序可生成递减序列。
	*/
	fmt.Println("After reversed: ", a)
}
