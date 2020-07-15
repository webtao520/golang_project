/*
   go context 包 WithTimeout 超时自动取消方法， 当执行一个go 协程，超时自动取消协程
*/

package main

import (
	"fmt"
	"context"
	"time"
)


// 模拟一个最小执行时间的阻塞函数
func inc(a int) int {
    res := a + 1                // 虽然我只做了一次简单的 +1 的运算,
    time.Sleep(1 * time.Second) // 但是由于我的机器指令集中没有这条指令,
    // 所以在我执行了 1000000000 条机器指令, 续了 1s 之后, 我才终于得到结果。B)
    return res
}
 
// 向外部提供的阻塞接口
// 计算 a + b, 注意 a, b 均不能为负
// 如果计算被中断, 则返回 -1
func Add(ctx context.Context, a, b int) int {
    res := 0
    for i := 0; i < a; i++ {
        res = inc(res)
        select {
        case <-ctx.Done():
            return -1
        default:
        }
    }
    for i := 0; i < b; i++ {
        res = inc(res)
        select {
        case <-ctx.Done():
            return -1
        default:
        }
    }
    return res
}

func main() {
    // 使用开放的 API 计算 a+b
    a := 1
    b := 2
	timeout := 2 * time.Second // 2s
    ctx, _ := context.WithTimeout(context.Background(), timeout)
    res := Add(ctx, 1, 2)
    fmt.Printf("Compute: %d+%d, result: %d\n", a, b, res)
}