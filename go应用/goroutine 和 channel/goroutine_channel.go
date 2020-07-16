package main

import (
	"fmt"
	"time"
)

var (
  myMap=make(map[int]int,10)
)

func test(n int){
	res:=1
	for i:=1;i<=n;i++ {
		res*=i
	}
	myMap[n]=res
}

func main (){
 // 开启多个协程完成这个任务[200个]
	for i:=1;i<=200;i++{
		go test(i)
	}
	// 休眠10秒钟 「都二个问题」
	time.Sleep(time.Second*10)
	//输出结果
	for i,v:=range myMap {
		fmt.Printf("map[%d]=%d\n",i,v)
	}
}