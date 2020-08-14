/*
		time.After
		time.After()表示多长时间长的时候后返回一条time.Time类型的通道消息。
		但是在取出channel内容之前不阻塞，后续程序可以继续执行。

	   先看源码（src/time/sleep.go）

		func After(d Duration) <-chan Time {
		return NewTimer(d).C
		}

		通过源码我们发现它返回的是一个NewTimer(d).C，其底层是用NewTimer实现的，
		所以如果考虑到效率低，可以直接自己调用NewTimer。

*/


package main

import (
   "fmt"
   "time"
)

func main() {
   t := time.After(time.Second * 3)
   fmt.Printf("t type=%T\n", t)
   //阻塞3秒
   fmt.Println("t=", <-t)
}

/*
		t type=<-chan time.Time
		t= 2020-08-14 14:57:31.7466332 +0800 CST m=+3.007249301
*/

