package main

import (
	"context"
	"fmt"
	"time"
)

// context.WithDeadline

/*
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
返回父上下文的副本，并将deadline调整为不迟于d。如果父上下文的deadline已经早于d，则WithDeadline(parent, d)在语义上等同于父上下文。
当截止日过期时，当调用返回的cancel函数时，或者当父上下文的Done通道关闭时，返回上下文的Done通道将被关闭，以最先发生的情况为准。
*/

func main() {
	d := time.Now().Add(2000 * time.Millisecond) // 2021-01-19 10:18:33.8327469 +0800 CST m=+2.002991901
	fmt.Println(d)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("周琳")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
