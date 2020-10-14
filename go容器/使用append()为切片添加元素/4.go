package main 

import (
	"fmt"
)

func main(){
	var a []int 
	a=append(a[:i],append([]int{x},a[i:]...)...) // 在第i个位置插入x
	a=append(a[:i],append([]int{1,2,3},a[i:]...)...)   // 在第i个位置插入切片
	fmt.Println(a)
}