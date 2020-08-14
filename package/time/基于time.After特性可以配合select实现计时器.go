/*
	基于time.After()特性可以配合select实现计时器
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 1
	for {
		select {
		case e1 := <-ch1:
			//如果ch1通道成功读取数据，则执行该case处理语句
			fmt.Printf("1th case is selected. e1=%v\n", e1)
		case <-time.After(time.Second * 2):
			fmt.Println("Timed out")
		}
	}

}

/*
	select语句阻塞等待最先返回数据的channel`,如ch1通道成功读取数据，
	则先输出1th case is selected. e1=1，之后每隔2s输出 Timed out。
*/
