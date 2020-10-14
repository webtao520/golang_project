package main 

import (
	"fmt"
)

func main (){
	var a [3]int // 定义三个整数的数组
	fmt.Println(a[0]) // 打印第一个元素 
	fmt.Println(a[len(a)-1]) // 打印最后一个元素
	// 打印索引和元素
	for i,v:=range a {
	   fmt.Printf("%d %d\n",i,v)
	}
   // 仅打印元素
   for _,v:=range a {
	   fmt.Printf("%d\n",v)
   }
}