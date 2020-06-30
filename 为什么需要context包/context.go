package main 


import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //创建等待组

func main (){
  // 需要开启几个goroutine就给等待组的计数器赋值为多少，这里是 2
  wg.Add(2)
  
}