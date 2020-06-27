package main 


import (
	"fmt"
)

/*
	select 语句类似于switch 语句，但是select会随机执行一个可运行的case，如果没有case可运行，它将阻塞，直到有case可运行
	
	select  分支语句，专门用户通道读写操作的
	  select {
		  case chan读/写：
			 分支1
		  case chan读/写：
			分支2
			...
		default:		 
	  }

	  执行流程：
			1. 每个case后的操作都会运算，如果有多个都可以执行，select 会随机执行一个可运行的case
			2. 如果没有case可运行。如果有default，执行default，如果没有default，它将阻塞，直到有case可运行
*/

func main (){
	  ch1:=make(chan int)
	  ch2:=make(chan int)

	  go func (){
     	ch1 <- 100
	  }()

	  go func(){
    	ch2 <- 200
	  }()
	  
	  select {
	  case data:= <- ch1:
		fmt.Println("ch1中读取数据了:",data)
	  case data:=<-ch2:
		fmt.Println("ch2中读取数据了:",data)
	  }

}