package main 

import (
	"fmt"
	"time"
)

func main()  {
	// 使用select 可以解决从管道取数据的阻塞问题
	//1.定义一个管道10个数据int
	intChan:=make(chan int, 10)
	for i:=0;i<10;i++ {
		intChan<-i
	}
	//2. 定义一个管道5个数据string
	stringChan:=make(chan string,5)
	for i:=0;i<5;i++ {
		stringChan<-"hello"+fmt.Sprintf("%d",i)
	}
	// 传统的方法再遍历管道时，如果不关闭会阻塞而导致deadLock
	// 问题，再实际开发中，可能我们不好确定什么关闭该管道
	//可以使用select方式可以解决
	for {
		select {
			// 注意：这里，如果intChan一直没有关闭，不会一直阻塞而deadlock
			// 会自动到下一个case匹配
		case v:=<-intChan:
			fmt.Printf("从intChan读取的数据%d\n",v)
			time.Sleep(time.Second)
		case v:=<-stringChan: 
			fmt.Printf("从stringChan读取的数据%s\n",v)
			time.Sleep(time.Second)
		default:
			fmt.Printf("都娶不到了，不玩了，程序员可以加入逻辑\n")
			time.Sleep(time.Second)
			return
		}
		
	}
}