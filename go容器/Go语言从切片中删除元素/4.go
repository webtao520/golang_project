package main 


import (
	"fmt"
)


func main(){
	a := []int{1, 2, 3}
	//a=append(a[:i],a[i+1]...)//  删除中间1个元素
	a=append(a[:1],a[2])
	fmt.Println(a)
}

/**

*a = []int{1, 2, 3, ...}
a = append(a[:i], a[i+1:]...) // 删除中间1个元素
a = append(a[:i], a[i+N:]...) // 删除中间N个元素
a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
a = a[:i+copy(a[i:], a[i+N:])] // 删除中间N个元素
*/