package  main 

import(
	"fmt"
	"time"
)

func main (){
	//  匿名函数创建goroutine
	go func(){
		var times int 
		for {
			 times ++ 
			 fmt.Println("tick",times)
			 time.Sleep(time.Second)
		}
	}()
	var input string 
	fmt.Scanln(&input)
}

/**
PS D:\goLang\github\golang_project\Go语言并发\Go语言轻量级线程> go run 2.go
tick 1
tick 2
tick 3
*/