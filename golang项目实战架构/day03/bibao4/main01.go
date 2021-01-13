package main 

import (
	"fmt"
)

func adder2(x int) func(int) int{
	return func (y int) int {
	   x+=y
	   return x
	}
  }
  
  
  // 闭包 = 函数 + 外部变量的引用
  func main(){
   ret := adder2(100)
   a:=ret
   fmt.Printf("%T\n",ret) // func(int) int
   fmt.Println(a)
  }