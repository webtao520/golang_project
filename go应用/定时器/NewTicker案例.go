package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	ch := make(chan bool)
	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Println("Ticker running .....")
			case stop := <-ch:
				if stop {
					fmt.Println("Ticker Stop")
					return
				}
			}
		}
	}(ticker)
	time.Sleep(5 * time.Second)
	ch <- true
	close(ch)
}

/*
func NewTicker

		func NewTicker(d Duration) *Ticker

		NewTicker返回一个新的Ticker，该Ticker包含一个通道字段，并会每隔时间段d就向该通道发送当时的时间。
		它会调整时间间隔或者丢弃tick信息以适应反应慢的接收者。如果d<=0会panic。关闭该Ticker可以释放相关资源。

*/
