package main 

import (
	"fmt"
)

func main(){
	a := []int{1,2,3}
	a=append(a[:1],a[1:]...)
	a=append(a[:0],a[1:]...)  // 删除开头1个元素
	a=append(a[:0],a[N:]...) // 删除开头N个元素
	fmt.Println(a)
}