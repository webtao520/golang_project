package main

import (
	"fmt"
)

// 常量
// 定义常量之后不能修改
// 在程序运行期间不会改变的量
const pi =  3.1415926

//  批量声明常量
const (
	statusOK = 200 
	notFound = 404
)

//  批量声明常量时，如果某一行声明后没有赋值，默认和上一行一致
const (
	n1 = 100
	n2 
	n3
)

// iota
const(
	a1 = iota   // 0
	a2   	   // 1
	a3        // 2
)

// 插队
const (
	c1 = iota // 0 
	c2 = 100 // 100
	c3 =iota  // 2
	c4 
)

// 多个常量声明在一行
const (
	d1,d2=iota+1,iota+2  // 0 // d1:1 d2:2
	d3,d4=iota+1,iota+2 // 1  // d3:2 d4:3
	d5 =iota +1 // 3
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)


func main (){
	// 非全局变量声明后必须使用
    fmt.Println(KB)
}


