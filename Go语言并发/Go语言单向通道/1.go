package main 

import (
	"fmt"
)

func main(){
	ch:=make(chan int)
	// 声明一个只能写入数据的通道类型，并赋值为ch
	var chSendOnly  chan<-int = ch
	// 声明一个只能读取数据的通道类型，并赋值为ch
	var chRecvOnly  <-chan int =ch
}

/**
上面的例子中，chSendOnly 只能写入数据，如果尝试读取数据，将会出现如下报错：
invalid operation: <-chSendOnly (receive from send-only type chan<- int)

同理，chRecvOnly 也是不能写入数据的。

*/