package main

import (
	"fmt"
)

func main(){
   var a []int 
   a=append(a,1)//追加1个元素
   a=append(a, 1,2,3) // 追加多个元素，手写解包方式
   a=append(a,[]int{1,2,3}...) // 追加一个切片，切片需要解包
   fmt.Println(a)
}

/**
PS D:\golang\github\golang_project\go容器\使用append()为切片添加元素> go run 1.go
[1 1 2 3 1 2 3]
*/