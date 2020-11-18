package main 

import (
	"fmt"
)

func main (){
 // c 保存包含10 的整型切片
 var c interface {} =[]int{10}
 // d 保存包含20的整型切片
 var d interface {}= []int{20}
 // 这里发生崩溃
 fmt.Println(c == d)	
}
