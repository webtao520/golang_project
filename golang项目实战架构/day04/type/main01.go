package main

import (
	"fmt"
)

// type  后面跟的是类型
type  myInt int // 自定义类型
type yourInt = int  // 类型别名

func main(){
   var  n myInt 
   n = 100
   fmt.Println(n)
   fmt.Printf("%T %d\n",n,n) // main.myInt 100

   var m yourInt
   m =100 
   fmt.Println(m)
   fmt.Printf("%T\n", m) // int

   var c  rune //  表示中文等非ASCll编码的字符
   c='中'
   fmt.Println(c) // 20013
   fmt.Printf("%T\n", c) // int32
}