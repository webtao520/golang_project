package main 


import (
	"fmt"
	"time"
)

/*
目前为止，我们仅仅演示了 Go 语言协程的启动和简单使用，但是通过上述代码还不足以验证协程是并发执行的，
接下来，我们通过下面这段代码来验证协程的并发执行：
*/

func add(a,b int){
	var c=a+b 
	fmt.Printf("%d + %d = %d\n",a,b,c)
	time.Sleep(8e9)
}

func main () {
   for  i:=0;i<10;i++ { // 启动了10个子协程
	   go add(1,i)
   }
   time.Sleep(1e9)
}