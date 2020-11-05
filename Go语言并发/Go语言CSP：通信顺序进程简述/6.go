
package main

import (
	"fmt"
	"sync"
)

func main(){
	// 声明一个同步等待组
	var wg  sync.WaitGroup  

	// 开N个后台打印线程
	for i:=0;i<10;i++ {
		wg.Add(1)
	
		go func (){
			fmt.Println("golang")
			//wg.Done()
			wg.Add(-1)
		}()
	}

	// 等待N个后台线程完成
	wg.Wait()
}


/**
PS D:\goLang\github\golang_project\Go语言并发\Go语言CSP：通信顺序进程简述> go run 6.go
golang
golang
golang
golang
golang
golang
golang
golang
golang
golang
*/