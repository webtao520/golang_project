package main 

import (
	"fmt"
	"time"
)

func main () {
	t:=time.After(time.Second *3)
	fmt.Printf("t type=%T\n", t)
	//阻塞3s
	fmt.Println("t=",<-t)
}

/*
time.After
	time.After()表示多长时间长的时候后返回一条time.Time类型的通道消息。但是在取出channel内容之前不阻塞，后续程序可以继续执行。

	先看源码

	func After(d Duration) <-chan Time {
	return NewTimer(d).C
	}

	通过源码我们发现它返回的是一个NewTimer(d).C，其底层是用NewTimer实现的，所以如果考虑到效率低，可以直接自己调用NewTimer。
*/


/*
PS D:\goLang\github\golang_project\go应用\定时器> go run .\time.After案例.go
					t type=<-chan time.Time
					t= 2020-07-09 14:57:49.9181999 +0800 CST m=+3.005762301
*/