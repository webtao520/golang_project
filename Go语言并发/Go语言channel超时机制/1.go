package  main 

import (
	"fmt"
	"time"
)

func main (){
	ch:=make(chan int)
	quit:=make(chan bool)
	
	// 新开一个协程
	go func (){
		for {
			select {
			case num:=<-ch:
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):
                fmt.Println("超时")
                quit <- true
              }	
			}
		}()

		for i:=0;i<5;i++{
			 ch <- i
			 time.Sleep(time.Second)
		}
		<-quit
		fmt.Println("程序结束")
	}

/**
PS D:\goLang\github\golang_project\Go语言并发\Go语言channel超时机制> go run 1.go
		num =  0
		num =  1
		num =  2
		num =  3
		num =  4
		超时
		程序结束
	*/