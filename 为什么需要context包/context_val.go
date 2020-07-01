package main

import (
    "fmt"
    "time"
    "context"
)

var key string = "name"

// 使用通过context向goroutinue传递值
func watch(ctx context.Context) {
    for {
        select{
        case <- ctx.Done():
            fmt.Println(ctx.Value(key), "退出，停止了。。。")
            return
        default:
            fmt.Println(ctx.Value(key), "运行中...")
            time.Sleep(2 * time.Second)
        }
    }
} 

func main() {
    ctx, cancel := context.WithCancel(context.Background())
  	// 给ctx绑定键值，传递给goroutine
    valuectx := context.WithValue(ctx, key, "【监控1】")
  	// 启动goroutine
    go watch(valuectx)

    time.Sleep(time.Second * 10)
    fmt.Println("该结束了。。。")
  	// 运行结束函数
    cancel()
    time.Sleep(time.Second * 3)
    fmt.Println("真的结束了。。")
}