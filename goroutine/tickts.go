package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	 练习题:
			模拟火车站买票
			火车票100张，4个售票口出售(4个goroutine)
*/
var wg sync.WaitGroup

func main() {
	wg.Add(4)
	go saleTickts("售票口1")
	go saleTickts("售票口2")
	go saleTickts("售票口3")
	go saleTickts("售票口4")
	wg.Wait()
	fmt.Println("改堂列车所有票卖光了....程序结束")
}

var tickts = 100

func saleTickts(name string) {
	rand.Seed(time.Now().UnixNano())
	for {
		if tickts > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)))
			fmt.Println(name, ":", tickts)
			tickts-- // 1 0 -1 -2 共享数据安全问题
		} else {
			fmt.Println(name, "结束买票....")
			break
		}
	}
	wg.Done()
}
