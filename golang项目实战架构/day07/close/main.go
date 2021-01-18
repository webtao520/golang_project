package main


import (
	"fmt"
)

// 关闭通道

func main (){
	 ch1:=make(chan bool,2) // 带缓冲区通道
	 ch1 <- true 
	 ch1 <-true 
	 close(ch1)

	//  for x:=range ch1{
	// 	fmt.Println(x) // 对一个关闭的通道进行接收会一直获取值直到通道为空。
	//  }

	// for {
	// 	x,ok :=<-ch1 
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(x)
	// }

	<-ch1
	<-ch1
	x, ok := <-ch1
	fmt.Println(x, ok)
	x, ok = <-ch1
	fmt.Println(x, ok)
	x, ok = <-ch1
	fmt.Println(x, ok)


}