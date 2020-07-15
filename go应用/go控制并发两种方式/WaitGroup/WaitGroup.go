package main

import (
	"fmt"
	"sync"
	//"time"
)

func main () {
	var wg sync.WaitGroup // 同步等待组
	wg.Add(2)
	go func(){
		//time.Sleep(2*time.Second)
		fmt.Println("1号完成")
		wg.Done()
	}()

	go func (){
		//time.Sleep(2*time.Second)
		fmt.Println("2号完成")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("好了，大家都干完了，放工")
}


/*
	PS D:\goLang\github\golang_project\go应用\go控制并发两种方式\WaitGroup> go run .\WaitGroup.go
							2号完成
							1号完成
							好了，大家都干完了，放工
*/