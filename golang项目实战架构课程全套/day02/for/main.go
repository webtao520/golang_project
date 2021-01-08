package main 


import (
	"fmt"
)

// 流程控制 跳出for 循环

func main(){
	//  当 i= 5 时 就跳出for循环
	for i:=0;i<10;i++ {
		if i == 5 {
			 break // 跳出 for 循环
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	// 当 i = 5 时， 跳过此次 for 循环，（不执行for循环内部的打印语句） 就绪下一次循环
	for i:=0;i<10;i++{
	 if i == 5 {
		 continue
	 }
	 fmt.Println(i)
	}
	fmt.Println("over2")
}