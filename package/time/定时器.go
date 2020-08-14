/*
		定时器（NewTicker）
*/

package main

import (
	 "fmt"
	 "time"
)

func main (){
	t:=time.NewTicker(time.Second*2)
	defer t.Stop()
	for {
		<-t.C 
		fmt.Println("Ticker running...")
	}
}



/*
			结果
			Ticker running…
			Ticker running…
			Ticker running…
			ticker只要定义完成后，不需要其他操作就可以定时执行。
			这里的defer t.Stop()和上面示例相似，也不会停止定时器，解决办法一样。
*/