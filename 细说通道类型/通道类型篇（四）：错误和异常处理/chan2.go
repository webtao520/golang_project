/*
	# 避免对已关闭通道进行操作 #

		为了避免对已关闭通道再度执行关闭操作引发 panic，一般我们约定只能在发送方关闭通道，而在接收方，
		我们则通过通道接收操作返回的第二个参数是否为 false 判定通道是否已经关闭，
		如果已经关闭，则不再执行发送操作，示例代码 如下：
*/

package main 


import (
	"fmt"
)

func main (){
   ch:=make(chan int,2)
   // 发送发
   go func (){
	   for i:=0;i<5;i++ {
		   fmt.Printf("发送发：发送数据 %v....\n",i)
		   ch<-i
	   }
	   fmt.Println("发送方：关闭通道。。。")
	   close(ch)
   }()

   // 接收方
   for {
	   num,ok:=<-ch 
	   if !ok{
		   fmt.Println("接收方： 通道已关闭")
		   break
	   }
	   fmt.Printf("接收方: 接收数据: %v\n", num)
   }
   fmt.Println("程序退出")
}

/*
		如果我们试图在通道 ch 关闭后发送数据到该通道，则会得到如下 panic：

		panic: send on closed channel

		而如果我们试图在通道 ch 关闭后再次关闭它，则会得到如下 panic：

		panic: close of closed channel
*/