package main 


import (
	"fmt"
)

// 浮点数

func main (){
   // math.MaxFloat32  float32 最大值
   f1:=1.23456 // float64
   fmt.Printf("%T\n",f1) // 默认go 语言中的小数都是float64 类型
   f2:=float32(1.23456) // 显示声明float32类型
   fmt.Printf("%T\n",f2) // float32
   // f1=f2 //    float32类型的值不能直接复赋值给float64类型的变量
   
}