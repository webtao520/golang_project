package main 


import (
	"fmt"
)

func main (){
	 ch3:=make(chan string, 5)
	 fmt.Printf("%T\n",ch3)
	 go sendData(ch3)
	 for  {
		 data,ok:=<-ch3
		 fmt.Println("\t读取数据 :",data)
		 if !ok {
			 fmt.Println("读取完毕....")
			 break
		 }
	 }
}

func sendData(ch3 chan string){
	for i:=1;i<=100;i++ {
		ch3 <- fmt.Sprint("数据 :",i)
		fmt.Println("已经写出数据 :", i)
	}
	close(ch3)
}