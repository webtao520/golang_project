package main 

import (
	"fmt"
)

func main (){
	//声明通道变量 var 通道变量 chan 通道类型 声明后需要配合 make 后才能使用。 chan 类型空值是 nil。
	ch1:=make(chan int) //  创建一个整型类型的通道
	ch2 := make(chan interface{})         // 创建一个空接口类型的通道, 可以存放任意格式
	type Equip  struct {/* 一些字段 */}
	ch2:=make(chan *Equip)  //  创建Equip指针类型的通道, 可以存放*Equip
}