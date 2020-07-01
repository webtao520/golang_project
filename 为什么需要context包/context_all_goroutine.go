package main

import (
    "fmt"
    "time"
    "context"
)

// 使用context控制多个goroutine
func watch(ctx context.Context, name string) {
    for {
        select {
        case <- ctx.Done():
            fmt.Println(name, "退出 ，停止了。。。")
            return
        default:
            fmt.Println(name, "运行中。。。")
            time.Sleep(2 * time.Second)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    go watch(ctx, "【任务1】")
    go watch(ctx, "【任务2】")
    go watch(ctx, "【任务3】")

    time.Sleep(time.Second * 10)
    fmt.Println("通知任务停止。。。。")
    cancel()
    time.Sleep(time.Second * 5)
    fmt.Println("真的停止了。。。")
}