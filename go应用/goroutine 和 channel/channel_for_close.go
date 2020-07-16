package main 

import (
	"fmt"
)

func main () {
  // 遍历管道
  intChan2:=make(chan int,100)
  for i:=0;i<100;i++ {
	  intChan2<-i*2 // 放入100个数据到管道
  }
  // 遍历时不能使用普通的for 循环
  /*
  for i:=0;i<len(intChan2);i++ {

  }
  */
  // 再遍历时，如果channel没有关闭，则会出现deadlock的错误
  // 再遍历时，如果channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历
  close(intChan2)
  for v:=range intChan2 {
	  fmt.Println("v=",v)
  }
}