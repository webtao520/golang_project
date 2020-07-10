package main

import (
	"fmt"
	"time"
)

func main() {
	// 再具体的一个时间点打印出golang
	myTicker := time.NewTicker(time.Second)
	for {
		nowTime := <-myTicker.C //  当前时间
		if nowTime.Hour() == 21 && nowTime.Minute() == 43 {
			fmt.Println("golang")
			break
		}
	}
}
