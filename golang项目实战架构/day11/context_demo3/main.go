package main 

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 为什么需要context？？
var wg sync.WaitGroup

func f2(ctx context.Context) {
	defer wg.Done()
FORLOOP:
	for {
		fmt.Println("保德路")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}
}

func f(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
FORLOOP:
	for {
		fmt.Println("周琳")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done(): // 只读的channel
			break FORLOOP
		default:
		}
	}
}


func main(){
	ctx, cancel := context.WithCancel(context.Background()) //根节点
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	// 如何通知子goroutine退出
	cancel()
	wg.Wait()
}