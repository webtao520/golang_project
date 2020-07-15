/*
	拥有超时控制的context有以下几种：

	context.WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) 指定时长超时结束
	context.WithCancel(parent Context) (ctx Context, cancel CancelFunc) 手动结束
	context.WithDeadline(parent Context, d time.Time) (Context, CancelFunc) 指定时间结束
	一般常用的话就context.WithTimeout
	以下示例代码
		所有超时控制结束的代码结构都是类似的，示例代码如下：
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*1)
	defer cancel() // 防止任务比超时时间短导致资源未释放
	// 启动协程
	go task(ctx)
	// 主协程需要等待，否则直接退出
	time.Sleep(time.Second * 4)
}

func task(ctx context.Context) {
	ch := make(chan struct{}, 0)
	// 真正的任务协程
	go func() {
		// 模拟两秒耗时任务
		time.Sleep(time.Second * 2)
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		fmt.Println("done")
	case <-ctx.Done(): // chan取消
		fmt.Println("timeout")
	}
}
