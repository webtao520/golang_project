package main 

import (
	"fmt"
)

// 定义一个结构体
type MyImplement  struct {

}

// 实现fmt.Stringer的String方法
func (m *MyImplement) String() string {
	 return "hi"
} 

// 在函数中返回fmt.Stringer接口
func  GetStringer () fmt.Stringer {
	var s *MyImplement =nil 
	 // 返回变量
	 return s
}


func main (){
	// 判断返回值是否为nil 
   if GetStringer () == nil {
	   fmt.Println("GetStringer() == nil")
   }else {
	   fmt.Println("GetStringer() !== nil")
   }
}